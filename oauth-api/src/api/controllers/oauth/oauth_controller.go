package oauth

import (
	"microservices-course/oauth-api/src/api/domain/oauth"
	"microservices-course/oauth-api/src/api/utils/apierrors"

	"github.com/gin-gonic/gin"
)

func CreateOauthToken(c *gin.Context) {
	var request oauth.AccessTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := apierrors.NewBadRequestApiError("Invalid payload")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

}
