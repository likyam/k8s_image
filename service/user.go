package service

import (
	"context"
)

var UserService = &userService{}

type userService struct {
}

func (u userService) GetUser(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	return &UserResponse{Username: "fafa1"}, nil
}

func (u userService) mustEmbedUnimplementedUserServer() {
	//TODO implement me
	panic("implement me")
}
