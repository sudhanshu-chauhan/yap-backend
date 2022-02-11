package app

import "time"

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	JWTToken string `json:"jwtToken"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JWTToken string `json:"jwtToken"`
}

type TaskRequest struct {
	Title  string    `json:"title"`
	Due    time.Time `json:"due"`
	Color  string    `json:"color"`
	Status string    `json:"status"`
}

type TaskResponse struct {
	Title  string    `json:"title"`
	Due    time.Time `json:"due"`
	Color  string    `json:"color"`
	Status string    `json:"status"`
	UserID int       `json:"userID"`
}
