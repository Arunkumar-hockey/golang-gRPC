package dto

import (
	"golang-gRPC/protomodel"
	"golang-gRPC/server/repository"
)

func NewUserCreateRequest(data *protomodel.User) *repository.User {
	return &repository.User{
		ID:       data.Id,
		Name:     data.Name,
		Email:    data.Email,
		PhoneNo:  data.Phoneno,
		Password: data.Password,
	}

}

func NewUserCreateResponse(data *repository.User) *protomodel.CreateUserRes {
	user := &protomodel.User{
		Id:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Phoneno:  data.PhoneNo,
		Password: data.Password,
	}
	return &protomodel.CreateUserRes{
		User: user,
	}

}
