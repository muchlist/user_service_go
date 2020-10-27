package services

import (
	"errors"
	"github.com/muchlist/erru_utils_go/logger"
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/user_service_go/domains/users"
	"github.com/muchlist/user_service_go/utils/crypt"
	"github.com/muchlist/user_service_go/utils/mjwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
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
	EditUser(string, users.UserEditRequest) (*users.UserResponse, rest_err.APIError)
	Login(users.UserLoginRequest) (*users.UserLoginResponse, rest_err.APIError)
	PutAvatar(email string, fileLocation string) (*users.UserResponse, rest_err.APIError)
}

//GetUser mendapatkan user dari domain
func (u *userService) GetUser(userID primitive.ObjectID) (*users.UserResponse, rest_err.APIError) {
	user, err := users.UserDao.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//FindUsers mendapatkan user dari domain
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

	hashPassword, errC := crypt.Obj.GenerateHash(user.Password)
	if errC != nil {
		logger.Error("Error pada kriptograpi", errC)
		return nil, rest_err.NewInternalServerError("Error pada kriptograpi", errors.New("bcrypt error"))
	}

	user.Password = hashPassword
	user.Timestamp = time.Now().Unix()

	insertedID, err := users.UserDao.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return insertedID, nil
}

func (u *userService) EditUser(userEmail string, request users.UserEditRequest) (*users.UserResponse, rest_err.APIError) {
	result, err := users.UserDao.EditUser(strings.ToLower(userEmail), request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) Login(login users.UserLoginRequest) (*users.UserLoginResponse, rest_err.APIError) {

	user, err := users.UserDao.GetUserByEmailWithPassword(login.Email)
	if err != nil {
		return nil, err
	}

	if !crypt.Obj.IsPWAndHashPWMatch(login.Password, user.HashPw) {
		return nil, rest_err.NewUnauthorizedError("Username atau password tidak valid")
	}

	claims := mjwt.CustomClaim{
		Identity:  user.Email,
		Name:      user.Name,
		IsAdmin:   user.IsAdmin,
		TimeExtra: 1,
		Jti:       "",
	}

	token, err := mjwt.Obj.GenerateToken(claims)
	if err != nil {
		return nil, err
	}

	userResponse := users.UserLoginResponse{
		Name:        user.Name,
		Email:       user.Email,
		IsAdmin:     user.IsAdmin,
		Avatar:      user.Avatar,
		AccessToken: token,
	}

	return &userResponse, nil

}

func (u *userService) PutAvatar(email string, fileLocation string) (*users.UserResponse, rest_err.APIError) {
	user, err := users.UserDao.PutAvatar(email, fileLocation)
	if err != nil {
		return nil, err
	}

	return user, nil
}
