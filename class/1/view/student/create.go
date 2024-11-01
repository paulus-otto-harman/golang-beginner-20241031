package student

import (
	handler "20241031/class/1/handler/student"
	userhandler "20241031/class/1/handler/user"
	"20241031/class/1/model"
	"20241031/class/1/util"
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type CreateStudent struct {
}

func (screen *CreateStudent) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Student Registration")

	name, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%s :", "Enter Student's name or [0] to exit")))))
	if name == "0" {
		return 0
	}
	username, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%s :", "Enter Student's username")))))

	tx, _ := db.Begin()

	util.BuildJson(model.User{
		Username: username,
		Password: "x",
		Active:   true,
		Role:     "Student",
	}, "body")
	userhandler.CreateUser(db)
	response := model.Response{}
	util.ReadJson(&response, "response")
	data := response.Data.(map[string]interface{})

	var userId int
	if value, ok := data["Id"]; ok {
		userId = int(value.(float64))
	}

	util.BuildJson(model.Student{
		Name:  name,
		User:  model.User{Id: userId},
		Batch: model.Batch{Id: 1},
	}, "body")
	handler.StudentRegistration(db)

	response = model.Response{}
	util.ReadJson(&response, "response")

	if response.StatusCode == 200 {
		tx.Commit()

		fmt.Println("Student successfully registered")
	}
	tx.Rollback()

	return util.GoBackOrNot()
}
