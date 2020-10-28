package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct lengkap dari document user di Mongodb
type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email" binding:"required,email"`
	Name      string             `json:"name" bson:"name"`
	IsAdmin   bool               `json:"is_admin" bson:"is_admin"`
	Avatar    string             `json:"avatar" bson:"avatar" binding:"required"`
	HashPw    string             `json:"hash_pw,omitempty" bson:"hash_pw,omitempty"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
}

//UserResponseList tipe slice dari UserResponse
type UserResponseList []UserResponse

//UserResponse struct kembalian dari MongoDB dengan menghilangkan hashPassword
type UserResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Name      string             `json:"name" bson:"name"`
	IsAdmin   bool               `json:"is_admin" bson:"is_admin"`
	Avatar    string             `json:"avatar" bson:"avatar"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
}

//UserRequest input JSON untuk keperluan register, timestamp dapat diabaikan
type UserRequest struct {
	Email     string `json:"email" bson:"email" binding:"required,email"`
	Name      string `json:"name" bson:"name" binding:"required"`
	IsAdmin   bool   `json:"is_admin" bson:"is_admin"`
	Avatar    string `json:"avatar" bson:"avatar"`
	Password  string `json:"password" bson:"password" binding:"required,min=3,max=20"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}

//UserEditRequest input JSON oleh admin untuk mengedit user
type UserEditRequest struct {
	Name            string `json:"name" bson:"name" binding:"required"`
	IsAdmin         bool   `json:"is_admin" bson:"is_admin"`
	TimestampFilter int64  `json:"timestamp_filter" bson:"timestamp" binding:"required"`
}

//UserLoginRequest input JSON oleh client untuk keperluan login
type UserLoginRequest struct {
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,min=3,max=20"`
}

//UserChangePasswordRequest struck untuk keperluan change password dan reset password
//pada reset password hanya menggunakan NewPassword dan mengabaikan Password
type UserChangePasswordRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password" binding:"required,min=3,max=20"`
	NewPassword string `json:"new_password" binding:"required,min=3,max=20"`
}

//UserLoginResponse balikan user ketika sukses login dengan tambahan AccessToken
type UserLoginResponse struct {
	Email       string `json:"email" bson:"email" binding:"required,email"`
	Name        string `json:"name" bson:"name" binding:"required"`
	IsAdmin     bool   `json:"is_admin" bson:"is_admin"`
	Avatar      string `json:"avatar" bson:"avatar"`
	AccessToken string `json:"access_token"`
}
