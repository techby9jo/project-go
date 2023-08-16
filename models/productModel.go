package models

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Productname string
	Detail      string
	Price       string
	Unit        string
}
