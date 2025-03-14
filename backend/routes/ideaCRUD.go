package routes

import "github.com/gin-gonic/gin"

func SetupIdeaRoutes(r *gin.Engine){
	r.GET("/ideas", controllers.getAllIdeas)
	r.GET("/idea/:id", controllers.getIdea)
	r.POST("/ideas/new", controllers.createNewIdea)
	r.DELETE("/idea/:id", controllers.deleteIdea)

}