package main

import (
	"gomenu/initializers"
	"gomenu/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Category{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Menu{})
	initializers.DB.AutoMigrate(&models.Endpoint{})
	initializers.DB.AutoMigrate(&models.Role{})
}
