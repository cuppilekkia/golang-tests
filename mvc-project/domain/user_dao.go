package domain

import (
	"microservices-course/mvc-project/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {
			ID:        123,
			FirstName: "Ciccio",
			LastName:  "Cappuccio",
			Email:     "via",
		},
	}

	UserDao UserDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type UserDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

//GetUser fetch the user
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, utils.NewApplicationError(
		"User not found",
		http.StatusNotFound,
		"not_found",
	)
}
