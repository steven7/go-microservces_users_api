package app

import (
	"github.com/gin-gonic/gin"
	"github.com/steven7/go-microservces_users_api/logger"
)

var (
	router = gin.Default()
)

func StartApplication()  {
	mapUrls()

	logger.GetLogger().Info("about to state the application...")
	router.Run(":8080")
}