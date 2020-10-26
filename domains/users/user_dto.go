package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct
type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email" binding:"required,email"`
	Name      string             `json:"name" bson:"name"`
	IsAdmin   bool               `json:"is_admin" bson:"is_admin"`
	Avatar    string             `json:"avatar" bson:"avatar" binding:"required"`
	HashPw    string             `json:"hash_pw,omitempty" bson:"hash_pw,omitempty"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
}

type UserResponseList []UserResponse

//UserResponse struct response
type UserResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Name      string             `json:"name" bson:"name"`
	IsAdmin   bool               `json:"is_admin" bson:"is_admin"`
	Avatar    string             `json:"avatar" bson:"avatar"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
}

//UserRequest input JSON oleh admin
type UserRequest struct {
	Email     string `json:"email" bson:"email" binding:"required,email"`
	Name      string `json:"name" bson:"name" binding:"required"`
	IsAdmin   bool   `json:"is_admin" bson:"is_admin"`
	Avatar    string `json:"avatar" bson:"avatar"`
	Password  string `json:"password" bson:"password" binding:"required,min=3,max=20"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}

//UserEditRequest input JSON oleh admin
type UserEditRequest struct {
	Name            string `json:"name" bson:"name" binding:"required"`
	IsAdmin         bool   `json:"is_admin" bson:"is_admin"`
	Avatar          string `json:"avatar" bson:"avatar"`
	TimestampFilter int64  `json:"timestamp_filter" bson:"timestamp" binding:"required"`
}

//UserLoginRequest input JSON oleh client
type UserLoginRequest struct {
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,min=3,max=20"`
}

//UserLoginResponse
type UserLoginResponse struct {
	Email   string `json:"email" bson:"email" binding:"required,email"`
	Name    string `json:"name" bson:"name" binding:"required"`
	IsAdmin bool   `json:"is_admin" bson:"is_admin"`
	Avatar  string `json:"avatar" bson:"avatar"`
	Token   string `json:"token"`
}

func (u *UserRequest) Validate() *error {
	return nil
}
