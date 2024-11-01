package view

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
)

type UpdateLesson struct {
}

func (screen *UpdateLesson) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Update Lesson")

	return util.GoBackOrNot()
}
