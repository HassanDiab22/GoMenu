package routes

import (
	"gomenu/controllers"
	"gomenu/dtos"
	"gomenu/middleware"

	"github.com/gin-gonic/gin"
)

func EndpointRoutes(r *gin.RouterGroup) {
	categoryRoutes := r.Group("/endpoint")
	{
		categoryRoutes.POST("/", middleware.ValidationMiddleware(dtos.CreateEndpointDTO{}), controllers.EndpointCategory)
		categoryRoutes.GET("/", controllers.GetAllCategories)
	}
}
