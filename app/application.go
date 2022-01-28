package app

import (
	"github.com/arun6783/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapURls()

	logger.Info("about to start our application")

	router.Run(":8080")
}
