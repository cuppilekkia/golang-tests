package oauth

import (
	"microservices-course/oauth-api/src/api/utils/apierrors"
	"strings"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *AccessTokenRequest) Validate() apierrors.ApiError {
	r.Username = strings.TrimSpace(r.Username)

	if r.Username == "" {
		return apierrors.NewBadRequestApiError("Invalid Username")
	}

	r.Password = strings.TrimSpace(r.Password)

	if r.Password == "" {
		return apierrors.NewBadRequestApiError("Invalid Password")
	}

	return nil
}
