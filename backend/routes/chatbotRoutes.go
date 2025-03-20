package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupChatbotRoutes(r *gin.Engine) {
	r.GET("/chat/new", controllers.GetIdeaFromChatbot)
}