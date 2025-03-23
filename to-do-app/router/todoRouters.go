package router

import (
	"to-do-app/controller"
	authMiddleWare "to-do-app/middleware"
	"to-do-app/models"

	"github.com/gin-gonic/gin"
)

func SetupToDoRouters(router *gin.Engine) {
	router.GET("/todo", controller.GetToDoTask) 
	router.POST("/todo", controller.CreateToDoTask)
	router.PUT("/todo/:id", controller.UpdateToDoTaskStatus)
	router.DELETE("/todo/:id", controller.DeleteToDoTask)
	router.PUT("/todo/assignTask/:id", authMiddleWare.AuthMiddlewareToken(models.Admin), controller.UpdateToDoTaskAssignee)
}
 
