package app

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB_HOST   string
	DB_PORT   int
	DB_USER   string
	DB_PASSWD string
	DB_NAME   string
}

func readConfig(path string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	return config, err
}

func GetConnection() *gorm.DB {
	config, configReadErr := readConfig("config.toml")
	if configReadErr != nil {
		panic("unable to read the config file")
	}
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASSWD,
		config.DB_NAME)
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
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
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
