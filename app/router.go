package app

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", Signup).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	return router
}
