package handler

import (
	"20241031/class/1/model"
	"20241031/class/1/repository"
	"20241031/class/1/service"
	"20241031/class/1/util"
	"database/sql"
)

func Login(db *sql.DB) {

	// input
	user := model.User{}
	util.ReadJson(&user, "body")

	// proses
	result := service.InitAuthService(*repository.CreateUser(db)).Login(user)

	// output
	util.BuildJson(result, "response")
}
