package view

import (
	"20241031/class/1/util"
	view "20241031/class/1/view/student"
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type Student struct {
}

func (screen *Student) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Students")
	fmt.Println("[1] Register")
	fmt.Println("[2] Delete")
	fmt.Println("[3] Update")
	fmt.Println("[4] View Students")

	fmt.Println()
	fmt.Println("[0] Return")

	menuItem, _ := gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[4] atau [0] untuk Kembali"))))
	switch menuItem.(int) {
	case 0:
		return 0
	case 1:
		Render(&view.CreateStudent{}, ctx, db)
	case 2:
		Render(&view.CreateStudent{}, ctx, db)
	case 3:
		Render(&view.CreateStudent{}, ctx, db)
	case 4:
		Render(&view.CreateStudent{}, ctx, db)
	}
	return -1
}
