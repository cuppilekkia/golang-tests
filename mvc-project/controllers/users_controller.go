package controllers

import (
	"microservices-course/mvc-project/services"
	"microservices-course/mvc-project/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		apiErr := utils.NewApplicationError(
			"ID format not valid",
			http.StatusBadRequest,
			"id_not_valid",
		)

		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userID)

	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
