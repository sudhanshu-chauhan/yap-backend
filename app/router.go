package app

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", Signup).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	// task endpoints
	router.HandleFunc("/task", CreateTask).Methods("POST")
	return router
}
