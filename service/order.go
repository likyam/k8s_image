package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var OrderService = &orderService{}

type orderService struct {
}

func (o orderService) GetOrder(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
	return &OrderResponse{
		OrderId:  10,
		UserName: o.getUserName(ctx),
	}, nil
}

func (o orderService) mustEmbedUnimplementedOrderServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (o orderService) getUserName(ctx context.Context) string {

	// 解析metada中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return ""
	}

	fmt.Println(md["x-request-id"])
	//header := metadata.New(map[string]string{
	//	"x-b3-traceid": md["x-request-id"][],
	//})
	//
	//var ctx = metadata.NewOutgoingContext(context.Background(), header)
	conn, err := grpc.Dial("user-server:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
