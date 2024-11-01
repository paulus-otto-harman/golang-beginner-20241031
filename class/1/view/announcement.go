package view

import (
	"20241031/class/1/util"
	view "20241031/class/1/view/announcement"
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type Announcement struct {
}

func (screen *Announcement) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Announcements")
	fmt.Println("[1] Create")
	fmt.Println("[2] Delete")
	fmt.Println("[3] Update")
	fmt.Println("[4] View Announcements")

	fmt.Println()
	fmt.Println("[0] Return")

	menuItem, _ := gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Select [1]-[4] or [0] to return"))))
	switch menuItem.(int) {
	case 0:
		return 0
	case 1:
		Render(&view.CreateAnnouncement{}, ctx, db)
	case 2:
		Render(&view.DeleteAnnouncement{}, ctx, db)
	case 3:
		Render(&view.UpdateAnnouncement{}, ctx, db)
	case 4:
		Render(&view.RetrieveAnnouncement{}, ctx, db)
	}
	return -1
}
