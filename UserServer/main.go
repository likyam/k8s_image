package main

import (
	"fmt"
	"google.golang.org/grpc"
	"istioDemo/service"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()

	service.RegisterUserServer(rpcServer, service.UserService)

	Listener, err := net.Listen("tcp", ":8003")
	if err != nil {
		panic(err)
	}
	err = rpcServer.Serve(Listener)
	if err != nil {
		panic(err)
	}
	fmt.Println("启动成功")
}
