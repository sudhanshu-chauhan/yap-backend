package app

import "golang.org/x/crypto/bcrypt"

func ResetPassword(email string, password string) bool {
	db := GetConnection()
	user := User{}
	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(password), 16)
	if hashErr != nil {
		return false
	}
	db.Model(&User{}).Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return false
	}
	db.Model(&user).Update("password_hash", passwordHash)
	return true
}
