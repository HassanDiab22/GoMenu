package main

import (
	"fmt"
	"gomenu/initializers"
	"gomenu/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	roles := []models.Role{
		{Name: "Admin"},
		{Name: "Client"},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := initializers.DB.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			initializers.DB.Create(&role)
			fmt.Printf("Role %s created\n", role.Name)
		} else {
			fmt.Printf("Role %s already exists\n", role.Name)
		}
	}
}
