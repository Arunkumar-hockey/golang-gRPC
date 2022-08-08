package service

import (
	"github.com/rs/xid"
	"golang-gRPC/server/repository"
)

var (
	UsersService UsersServiceInterface = &usersService{}
)

type UsersServiceInterface interface {
	Create(user *repository.User) (*repository.User, error)
}

type usersService struct{}

func (s *usersService) Create(user *repository.User) (*repository.User, error) {
	uid := xid.New()
	user.ID = uid.String()

	result, err := user.Create()
	if err != nil {
		return nil, err
	}
	return result, nil
}
