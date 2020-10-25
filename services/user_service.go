package services

import (
	"github.com/muchlist/user_service_go/domains/users"
	"github.com/muchlist/user_service_go/utils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// UserService publik
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(primitive.ObjectID) (*users.User, *errors.APIError)
	InsertUser(users.UserInput) (*string, *errors.APIError)
}

//GetUser mendapatkan user dari domain
func (u *userService) GetUser(userID primitive.ObjectID) (*users.User, *errors.APIError) {
	user, err := users.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) InsertUser(user users.UserInput) (*string, *errors.APIError) {
	insertedID, err := users.UserDao.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return insertedID, nil
}
