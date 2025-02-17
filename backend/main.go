package main

import (
	"log"

	"gomenu/controllers"
	"gomenu/initializers"
	"gomenu/middleware"
	"gomenu/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	// generateQRCode()

	r.POST("/signin", controllers.Signin)
	r.POST("/register", controllers.Register)

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())

	{
		routes.CategoryRoutes(authRoutes)
		routes.MenuRoutes(authRoutes)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	log.Println("Server running on :8080")
	r.Run()
}
