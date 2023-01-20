package models

import (
	"gorm.io/gorm"
)

type Burger struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Name        string
	Brand       string
	Description string
}
