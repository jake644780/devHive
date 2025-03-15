package controllers

import (
	"backend/db"
	"backend/models"
	"github.com/gin-gonic/gin"
)



func GetAllIdeas(c *gin.Context){
	var ideas []models.Idea
	result := db.Instance.Find(&ideas)
	if result.Error != nil{
		c.JSON(500, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, ideas)
}

func GetIdea(c *gin.Context){
	var idea models.Idea
	id := c.Param("id")
	result := db.Instance.First(&idea, id)
	if result.Error != nil{
		c.JSON(500, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, idea)

}

func CreateNewIdea(c *gin.Context){
	var idea models.Idea
	err := c.ShouldBindBodyWithJSON(&idea); if err != nil{
		c.JSON(500, gin.H{"error": err})
		return
	}
	result := db.Instance.Create(&idea); if result.Error != nil{
		c.JSON(500, gin.H{"error": result.Error})
		return
	}
	c.JSON(201, idea)
}

func DeleteIdea(c *gin.Context){
	var idea models.Idea
	id := c.Param("id")
	result := db.Instance.Delete(&idea, id)
	if result.Error != nil{
		c.JSON(500, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, idea)
}