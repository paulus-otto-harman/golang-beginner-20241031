package main

import (
	"20241031/class/1/handler"
	student_handler "20241031/class/1/handler/student"
	"20241031/database"
	"fmt"
)

func main() {
	db := database.DbOpen("20241028a")
	defer db.Close()

	var endpoint string
	fmt.Print("endpoint : ")
	fmt.Scan(&endpoint)

	switch endpoint {
	case "login":
		handler.Login(db)
	case "registrasi":
		student_handler.StudentRegistration(db)
	}
}

func m() {
	//wg := sync.WaitGroup{}
	//
	//wg.Add(1)
	//go app.Container(&wg, util.SessionTimeout)
	//wg.Wait()
}
