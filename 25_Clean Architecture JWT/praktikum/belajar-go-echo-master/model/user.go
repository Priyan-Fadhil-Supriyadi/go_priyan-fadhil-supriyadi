package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model

	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
