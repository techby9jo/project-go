package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string
	Lastname string
	Address  string
	Phone    string
	Username string
	Email    string
	Password string
	Role     string
}
