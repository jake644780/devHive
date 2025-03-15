package routes

import (
	"github.com/gin-gonic/gin"
	"backend/controllers"
)

func SetupIdeaRoutes(r *gin.Engine){
	r.GET("/ideas", controllers.GetAllIdeas)
	r.GET("/idea/:id", controllers.GetIdea)
	r.POST("/idea/new", controllers.CreateNewIdea)
	r.DELETE("/idea/:id", controllers.DeleteIdea)
}