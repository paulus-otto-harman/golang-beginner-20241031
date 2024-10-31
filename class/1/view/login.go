package view

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
	"errors"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type Login struct {
	Username string
}

func (login *Login) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Login")

	for err := errors.New(""); err != nil; {
		login.Username, err = gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-30s :", "Enter Username or [0] to exit")))))
	}

	if login.Username == "0" {
		return 0
	}

	password, _ := gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-30s :", "Password"))))

	if login.Username == "x" && password == "x" {
		menuUtama := Home{}
		Render(&menuUtama, ctx, db)
		if !menuUtama.isLogout {
			gola.Wait("Session expired. Press Enter to login")
		}
	}
	return -1
}
