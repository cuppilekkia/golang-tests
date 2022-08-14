package repositories

import (
	"microservices-course/src/api/domain/repositories"
	"microservices-course/src/api/logger"
	"microservices-course/src/api/services"
	"microservices-course/src/api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestApiError(err.Error())
		logger.Error("Error binding JSON", err)

		c.JSON(apiError.Status(), apiError)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}
