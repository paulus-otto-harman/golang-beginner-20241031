package student

import (
	"20241031/class/1/util"
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type CreateStudent struct {
}

func (screen *CreateStudent) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Student Registration")

	name, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%s :", "Enter Student's name or [0] to exit")))))
	if name == "0" {
		return 0
	}
	fmt.Println(name)

	return util.GoBackOrNot()
}
