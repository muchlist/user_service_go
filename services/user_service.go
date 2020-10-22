package services

import (
	"github.com/muchlist/user_service_go/domains"
	"github.com/muchlist/user_service_go/utils"
)

type userService struct{}

var (
	//UserService publik
	UserService userService
)

//GetUser mendapatkan user dari domain
func (u *userService) GetUser(userID int64) (*domains.User, *utils.ApplicationError) {
	user, err := domains.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, err
}
