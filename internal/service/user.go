package service

import "github.com/zhf0439/im-server/internal/repo"

type UserService struct {
	users *repo.UserRepo
}

func NewUserService(users *repo.UserRepo) *UserService {
	return &UserService{users: users}
}
