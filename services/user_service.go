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
	GetUser(primitive.ObjectID) (*users.UserResponse, rest_err.APIError)
	InsertUser(users.UserRequest) (*string, rest_err.APIError)
	FindUsers() (users.UserResponseList, rest_err.APIError)
}

//GetUser mendapatkan user dari domain
func (u *userService) GetUser(userID primitive.ObjectID) (*users.UserResponse, rest_err.APIError) {
	user, err := users.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//FindUsers mendapatkan users dari domain
func (u *userService) FindUsers() (users.UserResponseList, rest_err.APIError) {
	userList, err := users.UserDao.FindUser()
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (u *userService) InsertUser(user users.UserRequest) (*string, rest_err.APIError) {

	// cek ketersediaan email
	emailAvailable, err := users.UserDao.CheckEmailAvailable(user.Email)
	if err != nil {
		return nil, err
	}
	if !emailAvailable {
		return nil, rest_err.NewBadRequestError("Email tidak tersedia")
	}
	// END cek ketersediaan email

	insertedID, err := users.UserDao.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return insertedID, nil
}
