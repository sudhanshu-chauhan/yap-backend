package app

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(128);" json:"name"`
	Email string `gorm:"primaryKey;type:varchar(128);" json:"email"`
}

type AuthCred struct {
	gorm.Model
	UserEmail    string `gorm:"index" json:"userEmail"`
	Salt         string `gorm:"type:varchar(256);" json:"salt"`
	PasswordHash string `gorm:"type:varchar(512);" json:"passwordHash"`
	User         User   `gorm:"foreignkey:Email;references:UserEmail;"`
}
