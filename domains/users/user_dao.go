package users

import (
	"fmt"
	"github.com/muchlist/erru_utils_go/logger"
	"github.com/muchlist/erru_utils_go/rest_err"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"

	"github.com/muchlist/user_service_go/db"
)

const (
	connectTimeout = 2

	keyUserColl = "user"

	keyID        = "_id"
	keyEmail     = "email"
	keyHashPw    = "hash_pw"
	keyName      = "name"
	keyIsAdmin   = "is_admin"
	keyAvatar    = "avatar"
	keyTimeStamp = "timestamp"
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
	InsertUser(UserRequest) (*string, rest_err.APIError)
	GetUserByID(primitive.ObjectID) (*UserResponse, rest_err.APIError)
	GetUserByEmail(string) (*UserResponse, rest_err.APIError)
	GetUserByEmailWithPassword(string) (*User, rest_err.APIError)
	FindUser() (UserResponseList, rest_err.APIError)
	CheckEmailAvailable(string) (bool, rest_err.APIError)
	EditUser(string, UserEditRequest) (*UserResponse, rest_err.APIError)
}

//InsertUser menambahkan user
func (u *userDao) InsertUser(user UserRequest) (*string, rest_err.APIError) {

	coll := db.Db.Collection(keyUserColl)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	insertDoc := bson.D{
		{keyName, user.Name},
		{keyEmail, strings.ToLower(user.Email)},
		{keyIsAdmin, user.IsAdmin},
		{keyAvatar, user.Avatar},
		{keyHashPw, user.Password},
		{keyTimeStamp, user.Timestamp},
	}

	result, err := coll.InsertOne(ctx, insertDoc)
	if err != nil {
		apiErr := rest_err.NewInternalServerError("Gagal menyimpan user ke database", err)
		logger.Error("Gagal menyimpan user ke database", err)
		return nil, apiErr
	}

	insertID := result.InsertedID.(primitive.ObjectID).Hex()

	return &insertID, nil
}

//GetUser mendapatkan user dari database berdasarkan userID
func (u *userDao) GetUserByID(userID primitive.ObjectID) (*UserResponse, rest_err.APIError) {

	coll := db.Db.Collection(keyUserColl)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	var user UserResponse
	opts := options.FindOne()
	opts.SetProjection(bson.M{keyHashPw: 0})

	if err := coll.FindOne(ctx, bson.M{keyID: userID}, opts).Decode(&user); err != nil {

		if err == mongo.ErrNoDocuments {
			apiErr := rest_err.NewNotFoundError(fmt.Sprintf("User dengan ID %v tidak ditemukan", userID.Hex()))
			return nil, apiErr
		}

		logger.Error("gagal mendapatkan user (by ID) dari database", err)
		apiErr := rest_err.NewInternalServerError("Gagal mendapatkan user dari database", err)
		return nil, apiErr
	}

	return &user, nil
}

//GetUserByEmail mendapatkan user dari database berdasarkan email
func (u *userDao) GetUserByEmail(email string) (*UserResponse, rest_err.APIError) {

	coll := db.Db.Collection(keyUserColl)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	var user UserResponse
	opts := options.FindOne()
	opts.SetProjection(bson.M{keyHashPw: 0})

	if err := coll.FindOne(ctx, bson.M{keyEmail: strings.ToLower(email)}, opts).Decode(&user); err != nil {

		if err == mongo.ErrNoDocuments {
			apiErr := rest_err.NewNotFoundError(fmt.Sprintf("User dengan Email %s tidak ditemukan", email))
			return nil, apiErr
		}

		logger.Error("gagal mendapatkan user (by email) dari database", err)
		apiErr := rest_err.NewInternalServerError("Gagal mendapatkan user dari database", err)
		return nil, apiErr
	}

	return &user, nil
}

//GetUserByEmail mendapatkan user dari database berdasarkan email
func (u *userDao) GetUserByEmailWithPassword(email string) (*User, rest_err.APIError) {

	coll := db.Db.Collection(keyUserColl)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	var user User

	if err := coll.FindOne(ctx, bson.M{keyEmail: strings.ToLower(email)}).Decode(&user); err != nil {

		if err == mongo.ErrNoDocuments {
			// karena sudah pasti untuk keperluan login maka error yang dikembalikan unauthorized
			apiErr := rest_err.NewUnauthorizedError("Username atau password tidak valid")
			return nil, apiErr
		}

		logger.Error("Gagal mendapatkan user dari database (GetUserByEmailWithPassword)", err)
		apiErr := rest_err.NewInternalServerError("Gagal mendapatkan user dari database", err)
		return nil, apiErr
	}

	return &user, nil
}

//FindUser mendapatkan user dari database berdasarkan userID
func (u *userDao) FindUser() (UserResponseList, rest_err.APIError) {

	coll := db.Db.Collection(keyUserColl)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	users := UserResponseList{}
	opts := options.Find()
	opts.SetSort(bson.D{{keyID, -1}})
	sortCursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		logger.Error("Gagal mendapatkan user_handler dari database", err)
		apiErr := rest_err.NewInternalServerError("Database error", err)
		return UserResponseList{}, apiErr
	}

	if err = sortCursor.All(ctx, &users); err != nil {
		logger.Error("Gagal decode usersCursor ke objek slice", err)
		apiErr := rest_err.NewInternalServerError("Database error", err)
		return UserResponseList{}, apiErr
	}

	return users, nil
}

//CheckEmailAvailable melakukan pengecekan apakah alamat email sdh pernah ada di database
//jika ada akan return false ,yang artinya email tidak available
func (u *userDao) CheckEmailAvailable(email string) (bool, rest_err.APIError) {

	coll := db.Db.Collection(keyUserColl)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	opts := options.FindOne()
	opts.SetProjection(bson.M{keyID: 1})

	var user UserResponse

	if err := coll.FindOne(ctx, bson.M{keyEmail: strings.ToLower(email)}, opts).Decode(&user); err != nil {

		if err == mongo.ErrNoDocuments {
			return true, nil
		}

		logger.Error("Gagal mendapatkan user dari database", err)
		apiErr := rest_err.NewInternalServerError("Gagal mendapatkan user dari database", err)
		return false, apiErr
	}

	return false, nil
}

func (u *userDao) EditUser(userEmail string, userRequest UserEditRequest) (*UserResponse, rest_err.APIError) {
	coll := db.Db.Collection(keyUserColl)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(1)

	filter := bson.M{
		keyEmail:     userEmail,
		keyTimeStamp: userRequest.TimestampFilter,
	}
	update := bson.M{
		"$set": bson.M{
			keyName:      userRequest.Name,
			keyAvatar:    userRequest.Avatar,
			keyIsAdmin:   userRequest.IsAdmin,
			keyTimeStamp: time.Now().Unix(),
		},
	}

	var user UserResponse
	if err := coll.FindOneAndUpdate(ctx, filter, update, opts).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewBadRequestError("User tidak diupdate karena ID atau timestamp tidak valid")
		}

		logger.Error("Gagal mendapatkan user dari database", err)
		apiErr := rest_err.NewInternalServerError("Gagal mendapatkan user dari database", err)
		return nil, apiErr
	}

	return &user, nil
}
