package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint `gorm:"primary_key"`
	UID   string
	Iss   string
	Email string `gorm:"unique"`
}

type Credential struct {
	CredentialString string
}
