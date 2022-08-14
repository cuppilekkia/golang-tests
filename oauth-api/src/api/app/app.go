package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func StartApp() {
	router = gin.Default()

	mapUrls()

	log.Fatal(router.Run(":8081"))
}
