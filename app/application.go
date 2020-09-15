package app

import (
	"github.com/gin-gonic/gin"
	"go-microservces_users_api/logger"
)

var (
	router = gin.Default()
)

func StartApplication()  {
	mapUrls()

	logger.Log.Info("about to state the application...")
	router.Run(":8080")
}