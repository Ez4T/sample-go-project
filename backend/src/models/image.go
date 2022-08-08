package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Img_name string `json:"img_name"`
	Img_path string `json:"img_path"`
}
