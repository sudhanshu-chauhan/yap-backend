package main

import (
	"fmt"
	"yap-backend/app"
)

func main() {
	dbconn := app.GetConnection()
	dbconn.AutoMigrate(&app.User{})
	dbconn.AutoMigrate((&app.AuthCred{}))
	fmt.Println("database connection established")

}
