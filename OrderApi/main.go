package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"istioDemo/service"
)

func main() {

	fmt.Println("ok")

	r := gin.Default()
	// 添加 Jaeger 中间件
	//r.Use(Trace())

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

		header := metadata.New(map[string]string{
			"x-b3-traceid":      c.GetHeader("x-b3-traceid"),
			"x-b3-spanid":       c.GetHeader("x-b3-spanid"),
			"x-b3-parentspanid": c.GetHeader("x-b3-parentspanid"),
			"x-b3-sampled":      c.GetHeader("x-b3-sampled"),
			"x-b3-flags":        c.GetHeader("x-b3-flags"),
			"x-request-id":      c.GetHeader("x-request-id"),
			"test":              c.GetHeader("x-b3-traceid"),
			"test-x-request-id": c.GetHeader("x-request-id"),
		})

		var ctx = metadata.NewOutgoingContext(context.Background(), header)

		orderInfo, err := orderClient.GetOrder(ctx, request)

		fmt.Println(orderInfo)

		c.JSON(200, gin.H{
			"orderInfo":         orderInfo,
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
