package view

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
)

type DeleteLesson struct {
}

func (screen *DeleteLesson) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Delete a lesson")

	return util.GoBackOrNot()
}
