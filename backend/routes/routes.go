package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupAllRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "helo word",
		});
	});
	log.Println("loaded these routes:")
	SetupIdeaRoutes(r)
	SetupChatbotRoutes(r)
}