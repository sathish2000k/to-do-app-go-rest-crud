package controller

import (
	"log"
	"github.com/gin-gonic/gin"
	models "to-do-app/models"
	config "to-do-app/config"
	utils "to-do-app/utils"
	"time"
)

func GenerateToken(c *gin.Context) {
	log.Println("Generate Token")

	var user models.User
	var requestBody struct {
		Data struct {
			UserId int `json:"user_id"`
			Password string `json:"password"`
		} `json:"data"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := config.DB.Where("user_id = ?", requestBody.Data.UserId).First(&user).Error ;if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	log.Println("Fetched user: ", user)
	log.Println(user.Password, requestBody.Data.Password)

	if user.Password == "" {
		token, _ := utils.GenerateToken(user.UserId, string(user.UserRole), user.UserRole, time.Now().Add(15 * time.Minute))
		c.JSON(200, gin.H{
			"message": "Token generated successfully. Please reset the password within 15 Minutes",
			"token": token,
		})
		return 
	}

	err = utils.ComparePassword(user.Password, requestBody.Data.Password)
	if err != nil {
		c.JSON(403, gin.H{"error": "Password is incorrect"})
		return
	}

	token, _ := utils.GenerateToken(user.UserId, user.Password, user.UserRole,time.Now().Add(24 * time.Hour))

	if err != nil {
		c.JSON(400, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Token generated successfully",
		"token": token,
	})
}

func SetPassword(c *gin.Context) {
	log.Println("Set Password")

	var requestBody struct {
		Data struct {
			Password string `json:"password"`
			UserId int `json:"user_id"`
		} `json:"data"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	config.DB.Model(&models.User{}).Where("user_id = ?", requestBody.Data.UserId).First(&user)

	if user.Password == ""  {
		encrptedPassword := utils.HashPassword(requestBody.Data.Password)
		config.DB.Model(&models.User{}).Where("user_id = ?", requestBody.Data.UserId).Update("password", encrptedPassword)
		c.JSON(200, gin.H{
			"message": "Password set successfully",
			"user_id": requestBody.Data.UserId,
		})
	} else {
		c.JSON(400, gin.H{
			"error": "Password already set",
		})
		return
	}
}

func ResetPassword(c *gin.Context) {
	log.Println("Set Password")

	var requestBody struct {
		Data struct {
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
			UserId int `json:"user_id"`
		} `json:"data"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	config.DB.Model(&models.User{}).Where("user_id = ?", requestBody.Data.UserId).First(&user)

	err := utils.ComparePassword(user.Password, requestBody.Data.OldPassword)
	if err != nil {
		encrptedPassword := utils.HashPassword(requestBody.Data.NewPassword)
		config.DB.Model(&models.User{}).Where("user_id = ?", requestBody.Data.UserId).Update("password", encrptedPassword)
		c.JSON(200, gin.H{
			"message": "Password set successfully",
			"user_id": requestBody.Data.UserId,
		})
	} else  {
		c.JSON(400, gin.H{
			"error": "Old password is incorrect",
		})
		return 
	} 
}