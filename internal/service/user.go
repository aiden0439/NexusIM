package service

import "github.com/aiden0439/NexusIM/internal/repo"

type UserService struct {
	users *repo.UserRepo
}

func NewUserService(users *repo.UserRepo) *UserService {
	return &UserService{users: users}
}
