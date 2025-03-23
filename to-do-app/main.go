package main

import (
	"log"
	config "to-do-app/config"
	router "to-do-app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	routers := gin.Default()

	config.ConnectDB()
	config.ConnectRedis()
	
	router.SetupToDoRouters(routers)
	router.SetupAdminRouters(routers)
	router.SetupAuthRouters(routers)

	routers.Run(":8080")
	log.Println("Server started on port 8080")
}
