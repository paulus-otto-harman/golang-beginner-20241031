package view

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type CreateLesson struct {
}

func (screen *CreateLesson) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Add new lesson")

	lesson, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", "Lesson Name"))))
	fmt.Println(lesson)

	return util.GoBackOrNot()
}
