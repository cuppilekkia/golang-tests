package app

import (
	"microservices-course/src/api/controllers/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mapUrls() {
	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	router.POST("/repo", repositories.CreateRepo)
}
