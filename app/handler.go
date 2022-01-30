package app

import (
	"encoding/json"
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
