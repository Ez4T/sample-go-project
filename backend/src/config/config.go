package config

import (
	"api/web/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// local
	// dsn := "root:@tcp(127.0.0.1:3306)/interview?charset=utf8mb4&parseTime=True&loc=Local"
	// docker
	dsn := "root:interview@1234@tcp(db:3306)/interview?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Image{})
	DB = db
}
