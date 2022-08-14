package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

//StartApp starts the application
func StartApp() {
	mapUrls()

	log.Fatal(router.Run(":8080"))
}
