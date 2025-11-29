package services

import "lelang-online-api/repositories"

type UserService interface{}

type UserServiceImpl struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &UserServiceImpl{repo: repo}
}
