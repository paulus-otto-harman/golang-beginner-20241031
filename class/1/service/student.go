package service

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
)

type StudentService struct {
	Student repository.Student
}

func InitStudentService(repo repository.Student) *StudentService {
	return &StudentService{Student: repo}
}

func (repo *StudentService) Create(student model.Student) (*model.Response, error) {

	s, err := repo.Student.Create(&student)

	if err != nil {
		return nil, err
	}
	response := model.Response{
		StatusCode: 200,
		Message:    "Student Created",
		Data:       s,
	}
	return &response, nil
}
