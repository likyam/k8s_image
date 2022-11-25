package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"istioDemo/service"

	"net"
)

func main() {

	go healthz()

	rpcServer := grpc.NewServer()

	service.RegisterOrderServiceServer(rpcServer, service.OrderService)

	Listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	err = rpcServer.Serve(Listener)
	if err != nil {
		panic(err)
	}
	fmt.Println("启动成功")
}

func healthz() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
