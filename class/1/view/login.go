package view

import (
	handle "20241031/class/1/handler"
	"20241031/class/1/model"
	"20241031/class/1/util"
	com "20241031/class/1/util"
	"context"
	"database/sql"
	"errors"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type Login struct {
	Username string
	Password string
}

func (login *Login) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Login")

	for err := errors.New(""); err != nil; {
		login.Username, err = gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-30s :", "Enter Username or [0] to exit")))))
	}

	if login.Username == "0" {
		return 0
	}

	password, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-30s :", "Password")))))

	util.BuildJson(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{Username: login.Username, Password: password}, "body")
	handle.Login(db)
	response := model.Response{}
	util.ReadJson(&response, "response")

	switch response.StatusCode {
	case 200:
		home := Home{}
		Render(&home, ctx, db)
		if !home.isLogout {
			gola.Wait("Session expired. Press Enter to login")
		}
	case 401:
		gola.Wait(gola.Tf(gola.Color, "User Not Found", gola.LightYellow))
	}

	return -1
}
