package router

import (
	controller "to-do-app/controller"
	middleware "to-do-app/middleware"
	"to-do-app/models"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouters(router *gin.Engine) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/token", controller.GenerateToken)
		authRouter.PUT("/setPassword", middleware.AuthMiddlewareToken(models.Admin, models.Employee) , controller.SetPassword)
		authRouter.PUT("/resetPassword", middleware.AuthMiddlewareToken(models.Admin, models.Employee) , controller.ResetPassword)
	}
}