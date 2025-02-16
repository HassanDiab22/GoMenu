package controllers

import (
	"fmt"
	"gomenu/dtos"
	"gomenu/initializers"
	"gomenu/models"
	"gomenu/utils"
	"gomenu/validations"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	var requestData dtos.LoginDTO

	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.Validate(err, validations.SigninCustomValidationMessages, c)
		return
	}

	var user models.User
	if err := initializers.DB.Where("username= ?", requestData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username ",
		})
	}

	if !utils.CheckPassword(requestData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	userIDStr := strconv.Itoa(int(user.ID))
	token, err := utils.CreateToken(userIDStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func Register(c *gin.Context) {
	var requestData dtos.CreateUserDTO

	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.Validate(err, validations.RegisterCustomValidationMessages, c)
		return
	}

	var existingUser models.User
	if err := initializers.DB.Where("Username = ?", requestData.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}
	fmt.Println(existingUser)

	hashedPassword, err := utils.HashPassword(requestData.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Username: requestData.Username,
		Password: hashedPassword,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Registerd Successfully",
	})

}
