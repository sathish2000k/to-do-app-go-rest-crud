package controller

import (
	"log"
	config "to-do-app/config"
	models "to-do-app/models"

	"github.com/gin-gonic/gin"
)

func GetToDoTask(c *gin.Context) {
	log.Println("Get to-do task")

	var toDo []models.ToDo
	config.DB.Find(&toDo)

	c.JSON(200, gin.H{
		"message": "List of to-do tasks",
		"tasks": toDo,
	})
}

func GetToDoTaskById(c *gin.Context) {
	log.Println("Get to-do task by user id")

	user_id, _ := c.Get("user_id")

	var toDo models.ToDo
	config.DB.Model(&toDo).Where("task_assignee = ?", user_id).Find(&toDo)

	c.JSON(200, gin.H{
		"message": "Task details",
		"task": toDo,
	})
}

func CreateToDoTask(c *gin.Context) {
	var requestBody struct {
		Data struct {
			ToDo models.ToDo `json:"to_do"`
		} `json:"data"`
	}

	if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println("Create to-do task", requestBody.Data.ToDo)

	config.DB.Create(&requestBody.Data.ToDo)

	c.JSON(200, gin.H{
		"message": "Task created successfully",
		"task": requestBody.Data.ToDo,
	})
}

func UpdateToDoTaskStatus(c *gin.Context) {
	log.Println("Update to-do task")

	var requestBody struct {
		Data struct {
			TaskStatus string `json:"task_status"`
		}
	}

	taskId := c.Param("id")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&models.ToDo{}).Where("id = ?", taskId).Update("task_status", requestBody.Data.TaskStatus)

	c.JSON(200, gin.H{
		"message": "Task status updated successfully",
		"taskId": taskId,
		"taskStatus": requestBody.Data.TaskStatus,
	})
}

func DeleteToDoTask(c *gin.Context) {
	log.Println("Delete to-do task")

	taskId := c.Param("id")

	config.DB.Delete(&models.ToDo{}, taskId)

	c.JSON(200, gin.H{
		"message": "Task deleted successfully",
		"taskId": taskId,
	})
}

func UpdateToDoTaskAssignee(c *gin.Context) {
	taskId := c.Param("id")

	var requestBody struct {
		Data struct {
			AssigneeId int `json:"assignee_id"`
		}
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&models.ToDo{}).Where("id = ?", taskId).Update("task_assignee", requestBody.Data.AssigneeId)

	c.JSON(200, gin.H{
		"message": "Task assignee updated successfully",
		"taskId": taskId,
		"assigneeId": requestBody.Data.AssigneeId,
	})
}