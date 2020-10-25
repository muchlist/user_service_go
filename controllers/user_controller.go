package controllers

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/muchlist/user_service_go/domains"
	"github.com/muchlist/user_service_go/services"
	"github.com/muchlist/user_service_go/utils"

	"github.com/gin-gonic/gin"
)

//GetUser mengembalikan user
func GetUser(c *gin.Context) {
	//c.Param for parameter  localhost/users/user_id
	//c.Query for query localhost/users?user_id=
	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "UserID salah!",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		//c.JSON(apiErr.StatusCode, apiErr) <- normalnya seperti ini
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userID)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}

//InsertUser menambahkan user
func InsertUser(c *gin.Context) {

	var user domains.User
	if err := c.ShouldBindJSON(&user); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    fmt.Sprintf("Input tidak valid : %s", err.Error()),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	insertID, apiErr := services.UserService.InsertUser(user)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	res := gin.H{"msg": fmt.Sprintf("Register berhasil, ID: %s", *insertID)}

	utils.Respond(c, http.StatusOK, res)
}
