package view

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
)

type RetrieveLesson struct {
}

func (screen *RetrieveLesson) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Lessons")
	return util.WaitToGoBack()
}
