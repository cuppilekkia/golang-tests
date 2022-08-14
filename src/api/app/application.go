package app

import (
	"microservices-course/src/api/logger"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

//StartApp starts the application
func StartApp() {
	logger.Info("Starting APP", "myTag:app")
	mapUrls()

	logger.Log.Fatal(router.Run(":8080"))
}
