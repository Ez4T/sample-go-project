package controller

import (
	"api/web/config"
	"api/web/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DataReturn struct {
	Id        int
	Username  string
	Img_name  string
	Img_path  string
	Create_at string
}

type DataResponse struct {
	Images []DataReturn
}

const stringAsset string = "assets/"

func ValidateExt(value string) bool {
	if value == ".jpg" {
		return true
	}
	if value == ".jpeg" {
		return true
	}
	if value == ".png" {
		return true
	}
	return false
}

func (box *DataResponse) AddItem(item DataReturn) []DataReturn {
	box.Images = append(box.Images, item)
	return box.Images
}

func ListAllImages(c *gin.Context) {
	images := []models.Image{}
	result := []DataReturn{}
	box := DataResponse{result}
	config.DB.Find(&images)
	for _, image := range images {
		items := DataReturn{Id: image.Id, Username: image.Username, Img_path: stringAsset + image.Img_path, Img_name: image.Img_name, Create_at: image.CreatedAt.Local().String()}
		box.AddItem(items)
	}
	c.JSON(200, box)
}

func ListByUserImages(c *gin.Context) {
	images := []models.Image{}
	result := []DataReturn{}
	box := DataResponse{result}
	config.DB.Where("username = ?", c.Param("username")).Find(&images)
	for _, image := range images {
		items := DataReturn{Id: image.Id, Username: image.Username, Img_path: stringAsset + image.Img_path, Img_name: image.Img_name, Create_at: image.CreatedAt.Local().String()}
		box.AddItem(items)
	}
	c.JSON(200, box)
}

func ChangeImageName(c *gin.Context) {
	image := models.Image{}
	config.DB.Where("id = ?", c.Param("id")).First(&image)
	c.BindJSON(&image)
	config.DB.Save(&image)
	c.JSON(200, gin.H{"message": "Successfully changed image name."})
}

func DelImages(c *gin.Context) {
	image := models.Image{}
	config.DB.Where("id = ?", c.Param("id")).Delete(&image)
	c.JSON(200, gin.H{"message": "Successfully Deleted image."})
}

func Upload(c *gin.Context) {
	image := models.Image{}
	username := c.PostForm("username")
	file, err := c.FormFile("file")
	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	if !ValidateExt(extension) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid file type",
		})
		return
	}
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := username + "_" + uuid.New().String() + extension

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "./file/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	image.Username = username
	image.Img_name = newFileName
	image.Img_path = newFileName
	config.DB.Create(&image)
	// File saved successfully. Return proper result
	c.JSON(200, gin.H{"message": "Your file has been successfully uploaded."})
}

func MultiUpload(c *gin.Context) {
	image := models.Image{}
	username := c.PostForm("username")
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		extension := filepath.Ext(file.Filename)
		if !ValidateExt(extension) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid file type",
			})
			continue
		}

		newFileName := username + "_" + uuid.New().String() + extension
		if err := c.SaveUploadedFile(file, "./file/"+newFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			continue
		}

		image.Username = username
		image.Img_name = newFileName
		image.Img_path = newFileName
		config.DB.Create(&image)
		image.Id = image.Id + 1
	}

	// File saved successfully. Return proper result
	c.JSON(200, gin.H{"message": "Your file has been successfully uploaded."})
}
