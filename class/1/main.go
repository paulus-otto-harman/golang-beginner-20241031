package main

import (
	"20241031/class/1/app"
	"20241031/class/1/handler"
	studenthandler "20241031/class/1/handler/student"
	"20241031/class/1/util"
	"20241031/database"
	"database/sql"
	"fmt"
	"sync"
)

func main() {
	db := database.DbOpen("20241028a")
	defer db.Close()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go app.Container(&wg, util.SessionTimeout, db)
	wg.Wait()
}

func x(db *sql.DB) {
	var endpoint string
	fmt.Print("endpoint : ")
	fmt.Scan(&endpoint)

	switch endpoint {
	case "login":
		handler.Login(db)
	case "registrasi":
		studenthandler.StudentRegistration(db)
	}
}
