package routes

import (

	"github.com/gin-gonic/gin"
)

func SetupAllRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "dis is the backend, go to port 3000 for the frontend",
		});
	});
	SetupIdeaRoutes(r)
	SetupChatbotRoutes(r)
}