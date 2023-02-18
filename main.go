package main

import (
	"net/http"
	"yap-backend/app"
)

func main() {
	app.MigrateTables()
	router := app.GetRouter()
	http.ListenAndServe(":8000", router)

}
