package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"size:255;uniqueIndex"`
	Password string `json:"password"`
}
