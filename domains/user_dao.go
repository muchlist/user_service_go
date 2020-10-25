package domains

import (
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"

	"github.com/muchlist/user_service_go/db"
	"github.com/muchlist/user_service_go/utils"
)

const (
	connectTimeout = 2
)

var (
	//UserDao public
	UserDao userDaoInterface
)

func init() {
	//untuk keperluan testing sehingga dibuat interface
	UserDao = &userDao{}
}

type userDao struct {
}

type userDaoInterface interface {
	GetUser(primitive.ObjectID) (*User, *utils.ApplicationError)
	InsertUser(User) (*string, *utils.ApplicationError)
}

//InsertUser menambahkan user
func (u *userDao) InsertUser(user User) (*string, *utils.ApplicationError) {

	coll := db.Db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	insertDoc := bson.D{
		{Key: "name", Value: user.Name},
		{Key: "email", Value: user.Email},
		{Key: "is_admin", Value: user.IsAdmin},
		{Key: "avatar", Value: user.Avatar},
		{Key: "hash_pw", Value: user.HashPw},
	}

	result, err := coll.InsertOne(ctx, insertDoc)
	if err != nil {
		return nil, &utils.ApplicationError{
			Message:    "Gagal menyimpan user ke database",
			StatusCode: http.StatusInternalServerError,
			Code:       "mongo_db",
		}
	}

	insertID := result.InsertedID.(primitive.ObjectID).Hex()

	return &insertID, nil
}

//GetUser mendapatkan user dari database berdasarkan userID
func (u *userDao) GetUser(userID primitive.ObjectID) (*User, *utils.ApplicationError) {

	coll := db.Db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	var user User
	if err := coll.FindOne(ctx, bson.M{"_id": userID}).Decode(&user); err != nil {
		return nil, &utils.ApplicationError{
			Message:    "Gagal mendapatkan user dari database",
			StatusCode: http.StatusInternalServerError,
			Code:       "mongo_db",
		}
	}

	if user.Name == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("User %v tidak ditemukan", userID),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	return &user, nil
}
