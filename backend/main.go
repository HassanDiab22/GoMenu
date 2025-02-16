package main

import (
	"gomenu/controllers"
	"gomenu/initializers"
	"gomenu/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/signin", controllers.Signin)
	r.POST("/register", controllers.Register)

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())

	{
		authRoutes.POST("/category", controllers.CreateCategory)
		authRoutes.GET("/category", controllers.GetAllCategories)
		authRoutes.POST("/menu", controllers.CreatMenu)
		authRoutes.GET("/menu", controllers.GetAllMenus)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
