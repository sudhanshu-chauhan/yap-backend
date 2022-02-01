package app

import (
	"time"

	"gorm.io/gorm"
)

var dbconn = GetConnection()

type User struct {
	gorm.Model
	Name         string `gorm:"type:varchar(128);" json:"name"`
	Email        string `gorm:"primaryKey;type:varchar(128);" json:"email"`
	PasswordHash string `gorm:"type:varchar(256);" json:"passwordHash"`
	Type         string `gorm:"type:varchar(64);" json:"type"`
}

func (user User) create() error {
	return dbconn.Create(&user).Error
}

func (user User) getUser(email string) (User, error) {
	currentUser := User{}
	err := dbconn.First(&currentUser, "email = ?", email).Error
	return currentUser, err
}

type Task struct {
	Title  string    `gorm:"type:varchar(256)" json:"title"`
	UserID int       `gorm:"type:int64" json:"userId"`
	Due    time.Time `gorm:"type:timestamp;" json:"due"`
	Color  string    `gorm:"type:varchar(16);" json:"color"`
	Status string    `gorm:"type:varchar(24);" json:"status"`
}

func (task Task) create() error {
	return dbconn.Create(&task).Error
}
