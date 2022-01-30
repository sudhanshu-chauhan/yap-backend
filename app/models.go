package app

type User struct {
	Name     string   `gorm:"type:varchar(128);" json:"name"`
	Email    string   `gorm:"primaryKey;type:varchar(128);" json:"email"`
	AuthCred AuthCred `gorm:"foreignKey:PasswordHash" json:"authCred"`
}

type AuthCred struct {
	Salt         string `gorm:"type:varchar(256);" json:"salt"`
	PasswordHash string `gorm:"type:varchar(512);" json:"passwordHash"`
}
