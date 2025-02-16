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
}
