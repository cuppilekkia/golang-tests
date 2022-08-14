package services

import (
	"microservices-course/mvc-project/domain"
	"microservices-course/mvc-project/utils"
)

type userService struct{}

//UserService user services
var UserService userService

//GetUser get a single user
func (u *userService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
