package view

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type Home struct {
	isLogout bool
}

func (home *Home) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Home")
	fmt.Println("[1] Lessons")
	fmt.Println("[2] Announcements")

	fmt.Println()
	fmt.Println("[0] Logout")

	menuItem, _ := gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[5] atau [0] untuk Logout"))))
	switch menuItem.(int) {
	case 0:
		home.isLogout = true
		return 0
	case 1:
		Render(&Lesson{}, ctx, db)
	case 2:
		Render(&Announcement{}, ctx, db)
	}
	return -1
}
