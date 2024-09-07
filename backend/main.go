package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Image struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Filename string `gorm:"not null" json:"filename"`
}

var db *gorm.DB

func main() {
	var err error
	dsn := "root:1234@tcp(localhost:3306)/upload?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Image{})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/upload", uploadImage)
	r.GET("/images", getImages)
	r.GET("/images/:id", getImage)
	r.Static("/uploads", "./uploads")

	r.Run(":3000")
}

func uploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	filename := filepath.Base(file.Filename)
	savePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	image := Image{Filename: filename}
	result := db.Create(&image)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "id": image.ID})
}


func getImages(c *gin.Context) {
	var images []Image
	result := db.Find(&images)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	for i := range images {
		images[i].Filename = fmt.Sprintf("http://localhost:3000/uploads/%s", images[i].Filename)
	}

	c.JSON(http.StatusOK, images)
}

func getImage(c *gin.Context) {
	id := c.Param("id")
	var image Image
	result := db.First(&image, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	image.Filename = fmt.Sprintf("http://localhost:3000/uploads/%s", image.Filename)
	c.JSON(http.StatusOK, image)
}