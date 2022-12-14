package apierrors

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AError   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.AStatus
}
func (e *apiError) Message() string {
	return e.AMessage
}
func (e *apiError) Error() string {
	return e.AError
}

func NewApiError(status int, message string) ApiError {
	return &apiError{
		AStatus:  status,
		AMessage: message,
	}
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}

func NewBadRequestApiError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
