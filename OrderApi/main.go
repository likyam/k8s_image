package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"istioDemo/service"
)

func main() {

	fmt.Println("ok")

	r := gin.Default()
	r.Use(Trace())

	r.GET("/order", func(c *gin.Context) {
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
			"v":                 2,
			"x-request-id":      c.GetHeader("x-request-id"),
			"x-b3-traceid":      c.GetHeader("x-b3-traceid"),
			"x-b3-spanid":       c.GetHeader("x-b3-spanid"),
			"x-b3-parentspanid": c.GetHeader("x-b3-parentspanid"),
			"x-b3-sampled":      c.GetHeader("x-b3-sampled"),
			"x-b3-flags":        c.GetHeader("x-b3-flags"),
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
				LocalAgentHostPort: "istio-system.jaeger-collector:14268",
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
		//最开始的span，以url开始
		startSpan := tracer.StartSpan(ctx.Request.URL.Path)
		defer startSpan.Finish()
		ctx.Set("tracer", tracer)
		ctx.Set("parentSpan", startSpan)
		ctx.Next()
	}
}
