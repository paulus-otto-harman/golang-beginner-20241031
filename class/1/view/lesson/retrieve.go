package view

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
	gola "github.com/paulus-otto-harman/golang-module"
)

type RetrieveLesson struct {
}

func (screen *RetrieveLesson) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Lessons")
	gola.Wait("Press Enter to return to previous menu")
	return 0
}
