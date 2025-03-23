package router

import (
	"github.com/gin-gonic/gin"
	controller "to-do-app/controller"
	models "to-do-app/models"
	authMiddleWare "to-do-app/middleware"
)

func SetupAdminRouters(router *gin.Engine) {
	adminRouter := router.Group("/admin",  authMiddleWare.AuthMiddlewareToken(models.Admin), authMiddleWare.RateLimitMiddleWare())
	{
		adminRouter.POST("/createUser", controller.CreateUser)
		adminRouter.DELETE("/deleteUser/:id", controller.DeleteUser)
		adminRouter.GET("/getUser/:id", controller.GetUser)
		adminRouter.GET("/getAllUsers", controller.GetAllUsers)
	}
}