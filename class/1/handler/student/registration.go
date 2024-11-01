package handler_student

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
	"20241031/class/1/service"
	"20241031/class/1/util"
	"database/sql"
)

func StudentRegistration(db *sql.DB) {

	// input
	student := model.Student{}
	util.ReadJson(&student, "body")

	// proses
	result := service.InitStudentService(*repository.CreateStudent(db)).Create(student)

	// output
	util.BuildJson(result, "response")
}
