package app

import (
	"microservices-course/oauth-api/src/api/controllers/oauth"
	"microservices-course/oauth-api/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)

	router.POST("/oauth/access_token", oauth.CreateOauthToken)
}
