package utils

//ApplicationError standard application error struct
type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}

func NewApplicationError(mess string, status int, code string) *ApplicationError {
	return &ApplicationError{
		Message:    mess,
		StatusCode: status,
		Code:       code,
	}
}
