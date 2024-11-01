package handler

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
	"20241031/class/1/service"
	"20241031/class/1/util"
	"database/sql"
	"fmt"
)

func Login(db *sql.DB) {

	// input
	user := model.User{}
	util.ReadJson(&user, "body")

	// proses
	repoUser := repository.CreateUser(db)
	authentication := service.InitAuthService(*repoUser)

	result, err := authentication.Login(user)

	if err != nil {
		util.BuildJson(result, "response")
		fmt.Println("Login Handler :: Error : ", err)
	}

	// output
	util.BuildJson(result, "response")
}
