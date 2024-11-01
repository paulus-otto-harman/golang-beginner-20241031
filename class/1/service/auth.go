package service

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
)

type AuthService struct {
	User repository.User
}

func InitAuthService(repo repository.User) *UserService {
	return &UserService{User: repo}
}

func (repo *UserService) Login(user model.User) (*model.Response, error) {

	appUser, err := repo.User.Authenticate(&user)

	response := model.Response{}
	if err != nil {
		//return util.BuildResponse()
		response.StatusCode = 401
		response.Message = "login fail"
		response.Data = err.Error()
		return &response, err
	}
	response.StatusCode = 200
	response.Message = "login success"
	response.Data = appUser
	return &response, nil
}
