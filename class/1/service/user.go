package service

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
	"20241031/class/1/util"
)

type UserService struct {
	User repository.User
}

func InitUserService(repo repository.User) *UserService {
	return &UserService{User: repo}
}

func (repo *UserService) Create(user model.User) *model.Response {
	err := repo.User.Create(&user)

	if err != nil {
		return util.BuildResponse(401, "Fail to create user", struct {
			Message string
		}{
			Message: err.Error(),
		})
	}
	return util.BuildResponse(200, "User created", user)
}
