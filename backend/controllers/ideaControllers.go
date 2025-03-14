package controllers

import (
	"backend/db"
	"backend/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IdeaController(c *gin.Context, method string){
	db, err := db.InitDB()
	if err != nil{
		c.JSON(503, gin.H{"error":"couldn't open database"})
		return
	}
	switch method {
		case "all":
			GetAllIdeas(c, db)
		case "get":
			GetIdea(c, db)
		case "create":
			CreateNewIdea(c, db)
		case "delete":
			DeleteIdea(c, db)
	}
}



func GetAllIdeas(c *gin.Context, db *gorm.DB){
	var ideas []models.Idea
	result := db.Find(&ideas)
	if result.Error != nil{
		c.JSON(503, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, ideas)
}

func GetIdea(c *gin.Context, db *gorm.DB){
	var idea models.Idea
	id := c.Param("id")
	result := db.First(&idea, id)
	if result.Error != nil{
		c.JSON(503, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, idea)

}

func CreateNewIdea(c *gin.Context, db *gorm.DB){
	var idea models.Idea
	err := c.ShouldBindBodyWithJSON(&idea); if err != nil{
		c.JSON(500, gin.H{"error": err})
		return
	}
	result := db.Create(&idea); if result.Error != nil{
		c.JSON(500, gin.H{"error": result.Error})
		return
	}
	c.JSON(201, idea)
}

func DeleteIdea(c *gin.Context, db *gorm.DB){
	var idea models.Idea
	id := c.Param("id")
	result := db.Delete(&idea, id)
	if result.Error != nil{
		c.JSON(503, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, idea)
}