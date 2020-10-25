package services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/user_service_go/domains/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// UserService publik
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(primitive.ObjectID) (*users.User, *rest_err.APIError)
	InsertUser(users.UserInput) (*string, *rest_err.APIError)
}

//GetUser mendapatkan user dari domain
func (u *userService) GetUser(userID primitive.ObjectID) (*users.User, *rest_err.APIError) {
	user, err := users.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) InsertUser(user users.UserInput) (*string, *rest_err.APIError) {
	insertedID, err := users.UserDao.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return insertedID, nil
}
