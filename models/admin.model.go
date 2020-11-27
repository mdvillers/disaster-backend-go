package models

import (
	"gorm.io/gorm"
)

//Admin model
type Admin struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
