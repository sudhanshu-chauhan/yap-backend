package app

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	signupRequest := SignupRequest{}
	decodeErr := json.NewDecoder(r.Body).Decode(&signupRequest)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusInternalServerError)
	}

	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), 16)
	if hashErr != nil {
		http.Error(w, hashErr.Error(), http.StatusInternalServerError)
	}
	user := User{
		Name:         signupRequest.Name,
		Email:        signupRequest.Email,
		PasswordHash: string(passwordHash),
		Type:         "REGULAR",
	}

	err := user.create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, jwtErr := CreateJWTToken(signupRequest.Email)
	if jwtErr != nil {
		http.Error(w, jwtErr.Error(), http.StatusInternalServerError)
		return
	}
	signupResponse := SignupResponse{
		JWTToken: token,
	}
	res, marshalErr := json.Marshal(signupResponse)

	if marshalErr != nil {
		http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	loginReq := LoginRequest{}
	decodeErr := json.NewDecoder(r.Body).Decode(&loginReq)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusInternalServerError)
	}

	user, getUserError := User{}.getUser(loginReq.Email)
	if getUserError != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	hashCompareError := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password))
	if hashCompareError != nil {
		http.Error(w, "Wrong Credentials", http.StatusUnauthorized)
		return
	}
	token, jwtErr := CreateJWTToken(loginReq.Email)
	if jwtErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	signupResponse := SignupResponse{
		JWTToken: token,
	}

	res, marshalError := json.Marshal(signupResponse)
	if marshalError != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	email, err := DecodeJWTToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	currentUser, getUserErr := User{}.getUser(email)
	if getUserErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	taskRequest := TaskRequest{}
	log.Print(r.Body)
	requestDecodeError := json.NewDecoder(r.Body).Decode(&taskRequest)
	if requestDecodeError != nil {
		log.Print(requestDecodeError)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	task := Task{
		Title:  taskRequest.Title,
		Color:  taskRequest.Color,
		Status: taskRequest.Status,
		Due:    taskRequest.Due,
		UserID: int(currentUser.ID),
	}
	taskCreateError := task.create()
	if taskCreateError != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	res, marshalError := json.Marshal(TaskResponse{
		Title:  task.Title,
		Color:  task.Color,
		Status: task.Status,
		Due:    task.Due,
		UserID: task.UserID,
	})
	if marshalError != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func ListTask(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	email, err := DecodeJWTToken(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	currentUser, getUserErr := User{}.getUser(email)
	if getUserErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var tasks []Task
	GetTasks(int(currentUser.ID), &tasks)
	res, marshalError := json.Marshal(tasks)
	if marshalError != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
