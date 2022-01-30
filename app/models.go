package app

import (
	"time"
)

var dbconn = GetConnection()

type User struct {
	Name         string    `gorm:"type:varchar(128);" json:"name"`
	Email        string    `gorm:"primaryKey;type:varchar(128);" json:"email"`
	PasswordHash string    `gorm:"type:varchar(256);" json:"passwordHash"`
	Type         string    `gorm:"type:varchar(64);" json:"type"`
	CreatedAt    time.Time `gorm:"type:timestamp;autoCreateTime:nano;" json:"createdAt"`
	ModifiedAt   time.Time `gorm:"type:timestamp;autoCreateTime:nano;" json:"modifiedAt"`
}

func (user User) create() error {
	return dbconn.Create(&user).Error
}

func (user User) getUser(email string) (User, error) {
	currentUser := User{}
	err := dbconn.First(&currentUser, "email = ?", email).Error
	return currentUser, err
}
