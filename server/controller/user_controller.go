package controller

import (
	"context"
	"fmt"
	"golang-gRPC/protomodel"
	"golang-gRPC/server/dto"
	"golang-gRPC/server/service"
)

type userController struct{}

func (userController) mustEmbedUnimplementedUserServiceServer() {}

func NewUserControllerServer() UserServiceServer {
	return userController{}
}

func (userController) CreateUser(ctx context.Context, req *protomodel.CreateUserReq) (*protomodel.CreateUserRes, error) {
	fmt.Println("Create user request")
	user := dto.NewUserCreateRequest(req.GetUser())

	data, err := service.UsersService.Create(user)
	if err != nil {
		return nil, err
	}

	response := dto.NewUserCreateResponse(data)
	return response, nil
}
