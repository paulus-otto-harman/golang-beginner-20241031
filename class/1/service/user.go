package service

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
)

type UserService struct {
	User repository.User
}

func InitUserService(repo repository.User) *UserService {
	return &UserService{User: repo}
}

func (repo *UserService) LoginService(user model.User) (*model.Response, error) {

	customers, err := repo.User.Authenticate(&user)

	if err != nil {
		return nil, err
	}
	response := model.Response{
		StatusCode: 200,
		Message:    "login success",
		Data:       customers,
	}
	return &response, nil
}
