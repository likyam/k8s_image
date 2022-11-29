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
	r := gin.Default()
	r.GET("/order", func(c *gin.Context) {
		fmt.Println(c.GetHeader("x-request-id"))

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

		md := metadata.New(map[string]string{"x-request-id": c.GetHeader("x-request-id")})
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		orderInfo, err := orderClient.GetOrder(ctx, request)

		c.JSON(200, gin.H{
			"orderInfo": orderInfo,
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
