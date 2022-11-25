package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var OrderService = &orderService{}

type orderService struct {
}

func (o orderService) GetOrder(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
	println(o.getUserName())
	return &OrderResponse{
		OrderId:  10,
		UserName: o.getUserName(),
	}, nil
}

func (o orderService) mustEmbedUnimplementedOrderServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (o orderService) getUserName() string {
	conn, err := grpc.Dial("user_server:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	userClient := NewUserClient(conn)
	request := &UserRequest{Id: 1}
	userInfo, err := userClient.GetUser(context.Background(), request)
	fmt.Println(userInfo)
	return userInfo.GetUsername()
}
