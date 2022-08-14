package services

import (
	"microservices-course/mvc-project/domain"
	"microservices-course/mvc-project/utils"
	"net/http"
	"reflect"
	"testing"
)

var UserDaoMock userDaoMock

type userDaoMock struct{}

func (u *userDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	if userID == 123 {
		return &domain.User{
			ID:        123,
			FirstName: "a",
			LastName:  "b",
			Email:     "c",
		}, nil
	}

	return nil, &utils.ApplicationError{
		Message:    "User not found",
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}

func init() {
	domain.UserDao = &userDaoMock{}
}

func Test_userService_GetUser(t *testing.T) {
	type args struct {
		userID int64
	}
	tests := []struct {
		name  string
		u     *userService
		args  args
		want  *domain.User
		want1 *utils.ApplicationError
	}{
		{"User found", &UserService, args{123}, &domain.User{
			ID:        123,
			FirstName: "a",
			LastName:  "b",
			Email:     "c",
		}, nil},
		{"User not found", &UserService, args{1234}, nil, &utils.ApplicationError{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{}
			got, got1 := u.GetUser(tt.args.userID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("userService.GetUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
