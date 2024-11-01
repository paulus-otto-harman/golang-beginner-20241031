package student

import (
	handler "20241031/class/1/handler/student"
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
	fmt.Println(name)

	util.BuildJson(model.Student{
		Name:  name,
		User:  model.User{Id: 2},
		Batch: model.Batch{Id: 1},
	}, "body")

	handler.StudentRegistration(db)

	response := model.Response{}
	util.ReadJson(&response, "response")
	
	if response.StatusCode == 200 {
		fmt.Println("Student successfully registered")
	}

	return util.GoBackOrNot()
}
