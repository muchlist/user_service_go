package users

import (
	"errors"
	"fmt"
	"github.com/muchlist/erru_utils_go/logger"
	"github.com/muchlist/erru_utils_go/rest_err"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"

	"github.com/muchlist/user_service_go/db"
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
	GetUser(primitive.ObjectID) (*UserResponse, rest_err.APIError)
	InsertUser(input UserRequest) (*string, rest_err.APIError)
	FindUser() (*UserResponseList, rest_err.APIError)
}

//InsertUser menambahkan user
func (u *userDao) InsertUser(user UserRequest) (*string, rest_err.APIError) {

	coll := db.Db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	insertDoc := bson.D{
		{Key: "name", Value: user.Name},
		{Key: "email", Value: user.Email},
		{Key: "is_admin", Value: user.IsAdmin},
		{Key: "avatar", Value: user.Avatar},
		{Key: "hash_pw", Value: user.Password},
	}

	result, err := coll.InsertOne(ctx, insertDoc)
	if err != nil {
		apiErr := rest_err.NewInternalServerError("Gagal menyimpan user ke database", errors.New("database error"))
		logger.Error("Gagal menyimpan user ke database", err)
		return nil, apiErr
	}

	insertID := result.InsertedID.(primitive.ObjectID).Hex()

	return &insertID, nil
}

//GetUser mendapatkan user dari database berdasarkan userID
func (u *userDao) GetUser(userID primitive.ObjectID) (*UserResponse, rest_err.APIError) {

	coll := db.Db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	var user UserResponse
	opts := options.FindOne()
	opts.SetProjection(bson.M{"hash_pw": 0})

	if err := coll.FindOne(ctx, bson.M{"_id": userID}, opts).Decode(&user); err != nil {

		logger.Error("Gagal mendapatkan user dari database", err)
		apiErr := rest_err.NewInternalServerError("Gagal mendapatkan user dari database", errors.New("database error"))
		return nil, apiErr
	}

	if user.Name == "" {
		message := fmt.Sprintf("User %v tidak ditemukan", userID)
		logger.Info("Kembalian user dari database memiliki nama kosong")
		apiErr := rest_err.NewBadRequestError(message)
		return nil, apiErr
	}

	return &user, nil
}

//FindUser mendapatkan user dari database berdasarkan userID
func (u *userDao) FindUser() (*UserResponseList, rest_err.APIError) {

	coll := db.Db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	var users UserResponseList
	opts := options.Find()
	opts.SetSort(bson.D{{"_id", -1}})
	sortCursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		logger.Error("Gagal mendapatkan users dari database", err)
		apiErr := rest_err.NewInternalServerError("Database error", errors.New("database error"))
		return nil, apiErr
	}

	if err = sortCursor.All(ctx, &users); err != nil {
		logger.Error("Gagal decode usersCursor ke objek slice", err)
		apiErr := rest_err.NewInternalServerError("Database error", errors.New("database error"))
		return nil, apiErr
	}

	return &users, nil
}
