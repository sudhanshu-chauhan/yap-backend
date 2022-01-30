package app

import (
	"time"
)

var dbconn = GetConnection()

type User struct {
	Name       string    `gorm:"type:varchar(128);" json:"name"`
	Email      string    `gorm:"primaryKey;type:varchar(128);" json:"email"`
	CreatedAt  time.Time `gorm:"type:timestamp;autoCreateTime:nano;" json:"createdAt"`
	ModifiedAt time.Time `gorm:"type:timestamp;autoCreateTime:nano;" json:"modifiedAt"`
}

type AuthCred struct {
	Email        string    `gorm:"primaryKey;type:varchar(256);" json:"email"`
	PasswordHash string    `gorm:"type:varchar(512);" json:"passwordHash"`
	CreatedAt    time.Time `gorm:"type:timestamp;autoCreateTime:nano;" json:"createAt"`
	ModifiedAt   time.Time `gorm:"type:timestamp;autoCreateTime:nano" json:"modifiedAt"`
}

func (user User) create() error {
	return dbconn.Create(&user).Error
}

func (authCred AuthCred) create() error {
	return dbconn.Create(&authCred).Error
}
