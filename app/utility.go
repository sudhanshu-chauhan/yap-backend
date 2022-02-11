package app

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Uniqueharman10022020#"
	dbname   = "yap"
)

func GetConnection() *gorm.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type CustomClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateJWTToken(email string) (string, error) {
	claims := CustomClaim{
		email,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "YAP",
		},
	}
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := unsignedToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func DecodeJWTToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	customClaim, ok := token.Claims.(*CustomClaim)
	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return customClaim.Email, nil
}
