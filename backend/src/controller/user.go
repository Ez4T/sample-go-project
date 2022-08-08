package controller

import (
	"api/web/config"
	"api/web/lib"
	"api/web/models"
	"api/web/service"

	"github.com/gin-gonic/gin"
)

type BodyRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func GetUsers(c *gin.Context) {

	user := models.User{}
	json := BodyRequest{}
	c.BindJSON(&json)
	username := json.Username
	password := json.Password
	config.DB.Find(&user, "username = ?", username)
	decoded, err := lib.Decrypt(user.Password, MySecret)
	if err != nil {
		panic(err)
	}
	if decoded == password {
		token := service.JWTAuthService().GenerateToken(username, true)
		c.JSON(200, gin.H{"message": "success", "token": token})

	} else {
		c.JSON(200, gin.H{"message": "failed"})
	}
}

func CreateUsers(c *gin.Context) {
	user := models.User{}
	json := BodyRequest{}
	c.BindJSON(&json)
	if json.Username == "" {
		c.JSON(200, gin.H{"message": "error"})
		return
	}

	if json.Password == "" {
		c.JSON(200, gin.H{"message": "error"})
		return
	}
	password, err := lib.Encrypt(json.Password, MySecret)
	if err != nil {
		panic(err)
	}
	user.Username = json.Username
	user.Password = password
	result := config.DB.Create(&user)

	if result.Error == nil {
		c.JSON(200, gin.H{"message": "success"})
	} else {
		c.JSON(200, gin.H{"message": result.Error})
	}

}
