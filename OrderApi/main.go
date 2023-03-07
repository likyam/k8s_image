package main

import (
	"context"
	"fmt"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"istioDemo/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func main() {

	fmt.Println("ok")

	r := gin.Default()
	// 添加 Jaeger 中间件
	r.Use(Trace())

	r.GET("/order", func(c *gin.Context) {

		// 从请求上下文中获取跟踪上下文对象
		span := opentracing.SpanFromContext(c.Request.Context())
		if span == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// 设置跟踪标签
		span.SetTag("url", c.Request.URL.String())

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.Dial("order-server:8002", opts...)
		if err != nil {
			panic(err)
		}

		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)

		orderClient := service.NewOrderServiceClient(conn)

		request := &service.OrderRequest{Id: 1}

		//md := metadata.New(map[string]string{"x-request-id": c.GetHeader("x-request-id")})
		//ctx := metadata.NewOutgoingContext(context.Background(), md)
		//ctx := metadata.AppendToOutgoingContext(context.Background(), "x-request-id", c.GetHeader("x-request-id"), "x-request-id", c.GetHeader("x-request-id"),
		//	"x-b3-traceid", c.GetHeader("x-b3-traceid"),
		//	"x-b3-spanid", c.GetHeader("x-b3-spanid"),
		//	"x-b3-parentspanid", c.GetHeader("x-b3-parentspanid"),
		//	"x-b3-sampled", c.GetHeader("x-b3-sampled"),
		//	"x-b3-flags", c.GetHeader("x-b3-flags"))

		orderInfo, err := orderClient.GetOrder(context.WithValue(context.Background(), "ginContext", c), request)

		fmt.Println(orderInfo)

		c.JSON(200, gin.H{
			"orderInfo":         orderInfo,
			"v":                 3,
			"x-request-id":      c.GetHeader("x-request-id"),
			"x-b3-traceid":      c.GetHeader("x-b3-traceid"),
			"x-b3-spanid":       c.GetHeader("x-b3-spanid"),
			"x-b3-parentspanid": c.GetHeader("x-b3-parentspanid"),
			"x-b3-sampled":      c.GetHeader("x-b3-sampled"),
			"x-b3-flags":        c.GetHeader("x-b3-flags"),
			"X-Trace-ID":        c.Request.Header.Get("X-Trace-ID"),
		})
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}

// 拦截器
func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//jaeger配置
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1, //全部采样
			},
			Reporter: &jaegercfg.ReporterConfig{
				//当span发送到服务器时要不要打日志
				LogSpans:           true,
				LocalAgentHostPort: "jaeger-collector.istio-system.svc.cluster.local:14250",
			},
			ServiceName: "gin",
		}
		//创建jaeger
		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
		if err != nil {
			panic(err)
		}
		defer func(closer io.Closer) {
			err := closer.Close()
			if err != nil {

			}
		}(closer)
		// 设置全局tracer
		opentracing.SetGlobalTracer(tracer)

		// 从请求头中获取 Trace ID 和 Span ID
		traceID := ctx.Request.Header.Get("X-Trace-ID")
		spanID := ctx.Request.Header.Get("X-Span-ID")

		// 创建跟踪上下文对象
		var span opentracing.Span
		if traceID == "" || spanID == "" {
			span = tracer.StartSpan(ctx.Request.URL.Path)
		} else {
			parentSpanContext, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
			if err != nil {
				fmt.Printf("Failed to extract parent span: %v\n", err)
				ctx.Next()
				return
			}
			span = tracer.StartSpan(ctx.Request.URL.Path, opentracing.ChildOf(parentSpanContext))
		}

		// 将跟踪上下文对象存储到请求上下文中
		ctx.Request = ctx.Request.WithContext(opentracing.ContextWithSpan(ctx.Request.Context(), span))

		// 设置跟踪标签
		span.SetTag("method", ctx.Request.Method)
		span.SetTag("path", ctx.Request.URL.Path)

		// 将 Trace ID 和 Span ID 存储到响应头中
		ctx.Header("X-Trace-ID", span.Context().(jaeger.SpanContext).TraceID().String())
		ctx.Header("X-Span-ID", span.Context().(jaeger.SpanContext).SpanID().String())

		// 处理请求
		ctx.Next()

		// 结束跟踪
		span.Finish()
	}
}
