package view

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
)

type RetrieveAnnouncement struct {
}

func (screen *RetrieveAnnouncement) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Announcements")
	return util.WaitToGoBack()
}
