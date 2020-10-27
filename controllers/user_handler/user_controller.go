package user_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muchlist/erru_utils_go/logger"
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/user_service_go/domains/users"
	"github.com/muchlist/user_service_go/services"
	"github.com/muchlist/user_service_go/utils/mjwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"path/filepath"
)

//Get mengembalikan user
func Get(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		apiErr := rest_err.NewBadRequestError("Format userID salah")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userID)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

//Insert menambahkan user
func Insert(c *gin.Context) {

	var user users.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	insertID, apiErr := services.UserService.InsertUser(user)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	res := gin.H{"msg": fmt.Sprintf("Register berhasil, ID: %s", *insertID)}
	c.JSON(http.StatusOK, res)
}

//Find menampilkan list user
func Find(c *gin.Context) {

	userList, apiErr := services.UserService.FindUsers()
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": userList})
}

//Edit menampilkan list user
func Edit(c *gin.Context) {

	var user users.UserEditRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	userEdited, apiErr := services.UserService.EditUser(c.Param("user_email"), user)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, userEdited)
}

//Login
func Login(c *gin.Context) {

	var login users.UserLoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	response, apiErr := services.UserService.Login(login)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UploadImage(c *gin.Context) {

	claims := c.MustGet(mjwt.CLAIMS).(*mjwt.CustomClaim)

	file, err := c.FormFile("avatar")
	if err != nil {
		apiErr := rest_err.NewAPIError("File gagal di upload", http.StatusBadRequest, "bad_request", []interface{}{err.Error()})
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	fileName := file.Filename
	fileExtension := filepath.Ext(fileName)
	if !(fileExtension == ".jpg" || fileExtension == ".png" || fileExtension == ".jpeg") {
		apiErr := rest_err.NewBadRequestError("Ektensi file tidak di support")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	path := "static/images/" + claims.Identity + fileExtension

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		logger.Error(fmt.Sprintf("%s gagal mengupload file", claims.Identity), err)
		apiErr := rest_err.NewInternalServerError("File gagal di upload", err)
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	usersResult, apiErr := users.UserDao.PutAvatar(claims.Identity, path)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, usersResult)
}
