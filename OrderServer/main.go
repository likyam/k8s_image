package main

import (
	"fmt"
	"google.golang.org/grpc"
	"istioDemo/service"

	"net"
)

func main() {
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
