package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Token    string
	Role     string
}
