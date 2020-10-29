package services

import (
	"fmt"
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/user_service_go/domains/users"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"testing"
	"time"
)

var (
	getUserByIDFunction                func(userID primitive.ObjectID) (*users.UserResponse, rest_err.APIError)
	getUserByEmailFunction             func(email string) (*users.UserResponse, rest_err.APIError)
	getUserByEmailWithPasswordFunction func(email string) (*users.User, rest_err.APIError)
	insertUserFunction                 func(user users.UserRequest) (*string, rest_err.APIError)
	findUserFunction                   func() (users.UserResponseList, rest_err.APIError)
	checkEmailAvailableFunction        func(email string) (bool, rest_err.APIError)
	editUserFunction                   func(userEmail string, userRequest users.UserEditRequest) (*users.UserResponse, rest_err.APIError)
	deleteUserFunction                 func(userEmail string) rest_err.APIError
	putAvatarFunction                  func(email string, avatar string) (*users.UserResponse, rest_err.APIError)
	changePasswordFunction             func(data users.UserChangePasswordRequest) rest_err.APIError
)

func init() {
	//disini letak mockingnya, pada code asli interface di isi oleh Dao User
	//pada test di isi oleh usersDaoMock yang mengimplementasi semua Method interface
	users.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (u *usersDaoMock) GetUserByID(userID primitive.ObjectID) (*users.UserResponse, rest_err.APIError) {
	return getUserByIDFunction(userID)
}

func (u *usersDaoMock) GetUserByEmail(email string) (*users.UserResponse, rest_err.APIError) {
	return getUserByEmailFunction(email)
}

func (u *usersDaoMock) GetUserByEmailWithPassword(email string) (*users.User, rest_err.APIError) {
	return getUserByEmailWithPasswordFunction(email)
}

func (u *usersDaoMock) InsertUser(user users.UserRequest) (*string, rest_err.APIError) {
	return insertUserFunction(user)
}

func (u *usersDaoMock) FindUser() (users.UserResponseList, rest_err.APIError) {
	return findUserFunction()
}

func (u *usersDaoMock) CheckEmailAvailable(email string) (bool, rest_err.APIError) {
	return checkEmailAvailableFunction(email)
}

func (u *usersDaoMock) EditUser(userEmail string, userRequest users.UserEditRequest) (*users.UserResponse, rest_err.APIError) {
	return editUserFunction(userEmail, userRequest)
}

func (u *usersDaoMock) DeleteUser(userEmail string) rest_err.APIError {
	return deleteUserFunction(userEmail)
}

func (u *usersDaoMock) PutAvatar(email string, avatar string) (*users.UserResponse, rest_err.APIError) {
	return putAvatarFunction(email, avatar)
}

func (u *usersDaoMock) ChangePassword(data users.UserChangePasswordRequest) rest_err.APIError {
	return changePasswordFunction(data)
}

//--------------------------------------------------------------------------

func TestUserService_GetUser(t *testing.T) {
	getUserByIDFunction = func(userID primitive.ObjectID) (*users.UserResponse, rest_err.APIError) {
		return &users.UserResponse{
			ID:        primitive.NewObjectID(),
			Email:     "whois.muchlis@gmail.com",
			Name:      "Muchlis",
			IsAdmin:   true,
			Avatar:    "",
			Timestamp: time.Now().Unix(),
		}, nil
	}

	objectID := primitive.NewObjectID()
	user, err := UserService.GetUser(objectID)

	assert.Nil(t, err)
	assert.Equal(t, "Muchlis", user.Name)
	assert.Equal(t, "whois.muchlis@gmail.com", user.Email)
	assert.Equal(t, true, user.IsAdmin)
}

func TestUserService_GetUser_NoUserFound(t *testing.T) {
	getUserByIDFunction = func(userID primitive.ObjectID) (*users.UserResponse, rest_err.APIError) {
		apiErr := rest_err.NewNotFoundError(fmt.Sprintf("User dengan ID %v tidak ditemukan", userID.Hex()))
		return nil, apiErr
	}

	objectID := primitive.NewObjectID()
	user, err := UserService.GetUser(objectID)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("User dengan ID %v tidak ditemukan", objectID.Hex()), err.Message())
	assert.Equal(t, http.StatusNotFound, err.Status())
}

func TestUserService_GetUserByEmail_Found(t *testing.T) {
	getUserByEmailFunction = func(email string) (*users.UserResponse, rest_err.APIError) {
		return &users.UserResponse{
			ID:        primitive.NewObjectID(),
			Email:     "whois.muchlis@gmail.com",
			Name:      "Muchlis",
			IsAdmin:   true,
			Avatar:    "",
			Timestamp: time.Now().Unix(),
		}, nil
	}

	user, err := UserService.GetUserByEmail("whois.muchlis@gmail.com")

	assert.Nil(t, err)
	assert.Equal(t, "Muchlis", user.Name)
	assert.Equal(t, "whois.muchlis@gmail.com", user.Email)
	assert.Equal(t, true, user.IsAdmin)
}

func TestUserService_GetUserByEmail_NotFound(t *testing.T) {
	getUserByEmailFunction = func(email string) (*users.UserResponse, rest_err.APIError) {
		apiErr := rest_err.NewNotFoundError(fmt.Sprintf("User dengan Email %s tidak ditemukan", email))
		return nil, apiErr
	}

	user, err := UserService.GetUserByEmail("emailsembarang@gmail.com")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, "User dengan Email emailsembarang@gmail.com tidak ditemukan", err.Message())
	assert.Equal(t, http.StatusNotFound, err.Status())
}
