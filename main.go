package main

import (
	"net/http"
	"yap-backend/app"
)

func main() {

	router := app.GetRouter()
	http.ListenAndServe(":8000", router)

}
