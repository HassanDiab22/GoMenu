package routes

import (
	"gomenu/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(r *gin.RouterGroup) {
	menuRoutes := r.Group("/menu")
	{
		menuRoutes.POST("/", controllers.CreatMenu)
		menuRoutes.GET("/", controllers.GetAllMenus)
	}
}
