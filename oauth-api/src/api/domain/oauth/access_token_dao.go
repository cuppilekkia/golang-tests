package oauth

import (
	"microservices-course/mvc-project/domain"
	"microservices-course/oauth-api/src/api/utils/apierrors"
)

const (
	queryGetUserByUsernameAndPassword = ""
)

func GetUserByUsernameAndPassword(username string, password string) (*domain.User, apierrors.ApiError) {
	return nil, nil
}
