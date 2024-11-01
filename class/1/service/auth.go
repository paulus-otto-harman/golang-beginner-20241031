package service

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
	"20241031/class/1/util"
)

type AuthService struct {
	User repository.User
}

func InitAuthService(repo repository.User) *UserService {
	return &UserService{User: repo}
}

func (repo *UserService) Login(user model.User) *model.Response {
	appUser, err := repo.User.Authenticate(&user)

	if err != nil {
		return util.BuildResponse(401, "login fail", struct {
			Message string
		}{Message: err.Error()})
	}
	return util.BuildResponse(200, "login success", appUser)
}
