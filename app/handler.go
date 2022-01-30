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

	user := User{
		Name:  signupRequest.Name,
		Email: signupRequest.Email,
	}
	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), 16)
	if hashErr != nil {
		http.Error(w, hashErr.Error(), http.StatusInternalServerError)
	}

	authCred := AuthCred{
		Email:        signupRequest.Email,
		PasswordHash: string(passwordHash),
	}

	err := user.create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	authCredErr := authCred.create()
	if authCredErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	token, jwtErr := CreateJWTToken(signupRequest.Email)
	if jwtErr != nil {
		http.Error(w, jwtErr.Error(), http.StatusInternalServerError)
	}
	signupResponse := SignupResponse{
		JWTToken: token,
	}
	res, marshalErr := json.Marshal(signupResponse)

	if marshalErr != nil {
		http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func Login(w http.ResponseWriter, r *http.Request){
	loginReq := LoginRequest{}
	decodeErr := json.NewDecoder(r.Body).Decode(&loginReq)
}