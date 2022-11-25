package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"istioDemo/service"
)

func main() {
	r := gin.Default()
	r.GET("/order", func(c *gin.Context) {
		conn, err := grpc.Dial("order_server:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		orderInfo, err := orderClient.GetOrder(context.Background(), request)

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
