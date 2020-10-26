package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email" binding:"required,email"`
	Name     string             `json:"name" bson:"name"`
	IsAdmin  bool               `json:"is_admin" bson:"is_admin"`
	Avatar   string             `json:"avatar" bson:"avatar" binding:"required"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" binding:"required,min=3,max=20"`
	HashPw   string             `json:"hash_pw,omitempty" bson:"hash_pw,omitempty"`
}

type UserResponseList []UserResponse

//UserResponse struct response
type UserResponse struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Email   string             `json:"email" bson:"email"`
	Name    string             `json:"name" bson:"name"`
	IsAdmin bool               `json:"is_admin" bson:"is_admin"`
	Avatar  string             `json:"avatar" bson:"avatar"`
}

//UserRequest input JSON oleh admin
type UserRequest struct {
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Name     string `json:"name" bson:"name" binding:"required"`
	IsAdmin  bool   `json:"is_admin" bson:"is_admin"`
	Avatar   string `json:"avatar" bson:"avatar"`
	Password string `json:"password" bson:"password" binding:"required,min=3,max=20"`
}

func (u *UserRequest) Validate() *error {
	return nil
}
