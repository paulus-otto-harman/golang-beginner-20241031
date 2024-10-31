package handler

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

func Login(db *sql.DB) {

	// input
	user := model.User{}
	file, err := os.Open("body.json")

	if err != nil {
		fmt.Println("Login Error : ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&user)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
	}

	// proses
	repoUser := repository.CreateUser(db)
	customerService := service.InitUserService(*repoUser)

	result, err := customerService.LoginService(user)

	if err != nil {
		fmt.Println("Login 2 :: Error : ", err)
	}

	// output
	jsonData, err := json.MarshalIndent(result, " ", "")

	if err != nil {
		fmt.Println("err :", err)
	}

	fmt.Println(string(jsonData))
}
