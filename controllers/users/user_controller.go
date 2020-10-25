package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muchlist/user_service_go/domains/users"
	"github.com/muchlist/user_service_go/services"
	"github.com/muchlist/user_service_go/utils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Get mengembalikan user
func Get(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		apiErr := errors.NewBadRequestError("Format userID salah")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userID)
	if apiErr != nil {
		c.JSON(http.StatusNotFound, apiErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

//Insert menambahkan user
func Insert(c *gin.Context) {

	var user users.UserInput
	if err := c.ShouldBindJSON(&user); err != nil {
		apiErr := errors.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	insertID, apiErr := services.UserService.InsertUser(user)
	if apiErr != nil {
		c.JSON(http.StatusNotModified, apiErr)
		return
	}

	res := gin.H{"msg": fmt.Sprintf("Register berhasil, ID: %s", *insertID)}
	c.JSON(http.StatusOK, res)
}
