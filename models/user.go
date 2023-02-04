package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint `gorm:"primary_key"`
	Name  string
	Iss   string
	Email string `gorm:"unique"`
}

type Credential struct {
	CredentialString string
}
