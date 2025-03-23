package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	Admin Role = "Admin"
	Employee Role = "User"
)

type User struct {
	gorm.Model
	UserId int `json:"user_id" gorm:"primaryKey"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	UserRole Role `json:"role"`
}