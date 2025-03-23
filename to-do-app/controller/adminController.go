package controller

import (
	"log"
	"github.com/gin-gonic/gin"
	models "to-do-app/models"
	config "to-do-app/config"
)

func GetUser(c *gin.Context)  {
	log.Println("Get User")

	var user models.User

	userId := c.Param("id")
	config.DB.First(&user, userId)

	c.JSON(200, gin.H{
		"message": "User details",
		"user": user,
	})
}

func CreateUser(c *gin.Context) {
	log.Println("Create User")

	var requestBody struct {
		Data struct {
			User models.User `json:"user"`
		}
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user = requestBody.Data.User

	config.DB.Create(&user)

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	log.Println("Delete User")

	userId := c.Param("id")

	config.DB.Delete(&models.User{}, userId)

	c.JSON(200, gin.H{
		"message": "User deleted successfully",
		"userId": userId,
	})
}

func GetAllUsers(c *gin.Context) {
	log.Println("Get All Users")

	var users []models.User

	config.DB.Select("user_id, user_name, user_role").Find(&users)

	c.JSON(200, gin.H{
		"message": "List of all users",
		"users": users,
	})
}