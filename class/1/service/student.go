package service

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
	"20241031/class/1/util"
)

type StudentService struct {
	Student repository.Student
}

func InitStudentService(repo repository.Student) *StudentService {
	return &StudentService{Student: repo}
}

func (repo *StudentService) Create(student model.Student) *model.Response {

	s, err := repo.Student.Create(&student)

	if err != nil {
		return util.BuildResponse(401, "Fail to register new Student", struct {
			Message string
		}{
			Message: err.Error(),
		})
	}
	return util.BuildResponse(200, "Student Created", s)

}
