package domains

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct
type User struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email   string             `json:"email" bson:"email" binding:"required,email"`
	Name    string             `json:"name" bson:"name"`
	IsAdmin bool               `json:"is_admin" bson:"is_admin"`
	Avatar  string             `json:"avatar" bson:"avatar" binding:"required"`
	Pw      string             `json:"pw,omitempty" bson"pw,omitempty" binding:"required"`
	HashPw  string             `json:"hash_pw,omitempty" bson:"hash_pw,omitempty"`
}
