package services

import (
	"microservices-course/oauth-api/src/api/domain/oauth"
	"microservices-course/oauth-api/src/api/utils/apierrors"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, apierrors.ApiError)
}

var OauthService oauthServiceInterface

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, apierrors.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}
