package app

import (
	"20241031/class/1/view"
	"20241031/database"
	"context"
	"sync"
	"time"
)

func sessionTimeout(timeout int) time.Duration {
	if timeout == 0 {
		return time.Until(time.Now().AddDate(10, 0, 0))
	}
	return time.Duration(timeout) * time.Second
}

func Container(wg *sync.WaitGroup, timeout int) {
	db := database.DbOpen("20241028a")
	defer db.Close()

	defer wg.Done()
	sessionLifetime := sessionTimeout(timeout)
	for {
		loginScreen := view.Login{}
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), sessionLifetime)
			defer cancel()

			view.Render(&loginScreen, ctx, db)
		}()
		if loginScreen.Username == "0" {
			return
		}
	}
}