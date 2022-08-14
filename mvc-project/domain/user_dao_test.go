package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(1234)

	if user != nil {
		t.Error("A wrong ID should not return a User")
	}

	if err == nil {
		t.Error("A wrong ID should return an error")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("A wrong ID should return a StatusNotFound code")
	}
}

func TestGetUserFound(t *testing.T) {
	user, err := UserDao.GetUser(123)

	if user == nil {
		t.Error("A correct ID should return a User")
	}

	if err != nil {
		t.Error("A correct ID should not return an error")
	}
}
