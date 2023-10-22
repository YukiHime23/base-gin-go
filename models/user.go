package models

import (
	"base-gin-go/pkg/hash"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required,usernameAllow"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required,passwordAllow" swaggerignore:"true"`
} //@name User

func (u *User) HashPassword(password string) {
	u.Password = hash.GenerateFromPassword(password)
}

type RequestRegister struct {
	Username string `json:"username" binding:"required,usernameAllow"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
} //@name RequestRegister

type RequestLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
} //@name RequestLogin
