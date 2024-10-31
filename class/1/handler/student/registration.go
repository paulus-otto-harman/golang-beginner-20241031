package handler_student

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
	"20241031/class/1/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func StudentRegistration(db *sql.DB) {

	// input
	student := model.Student{}
	file, err := os.Open("student.body.json")

	if err != nil {
		fmt.Println("Student JSON Error : ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&student)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
	}

	// proses
	repoStudent := repository.CreateStudent(db)
	studentService := service.InitStudentService(*repoStudent)

	result, err := studentService.Create(student)

	if err != nil {
		fmt.Println("Student Repo :: Error : ", err)
	}

	// output
	jsonData, err := json.MarshalIndent(result, " ", "")

	if err != nil {
		fmt.Println("err :", err)
	}

	fmt.Println(string(jsonData))
}
