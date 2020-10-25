package services

import (
	"github.com/muchlist/user_service_go/domains"
	"github.com/muchlist/user_service_go/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct{}

var (
	//UserService publik
	UserService userService
)

//GetUser mendapatkan user dari domain
func (u *userService) GetUser(userID primitive.ObjectID) (*domains.User, *utils.ApplicationError) {
	user, err := domains.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *userService) InsertUser(user domains.User) (*string, *utils.ApplicationError) {
	insertedID, err := domains.UserDao.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return insertedID, err
}
