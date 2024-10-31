package view

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
	gola "github.com/paulus-otto-harman/golang-module"
)

type RetrieveAnnouncement struct {
}

func (screen *RetrieveAnnouncement) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Announcements")
	gola.Wait("Press Enter to return to previous menu")
	return 0
}
