package controllers

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StringifyIdeas(ideas []models.Idea) string{
	var buffer string
	for i, idea := range ideas{
		buffer += fmt.Sprintf(`
		%v. idea:
		\n\tcategory: %v 
		\n\tContent: %v
		`, i + 1, idea.Category, idea.Content)
	}
	return buffer
}

func GetIdeaFromChatbot(c *gin.Context){
	// Call the /getAllIdeas route on the Gin server
	resp, err := http.Get("http://localhost:8081/ideas")
	if err != nil{
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("%v", err.Error())

	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("%v", err.Error())
	}
	// Parse the JSON response into a slice of Idea
	var ideas []models.Idea
	if err := json.Unmarshal(body, &ideas); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("%v", err.Error())
	}

	prompt := StringifyIdeas(ideas)
	if err != nil{
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		log.Fatalf("%v", err.Error())
	}
	c.String(200, prompt)

	
}
