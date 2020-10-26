package app

import (
	"github.com/muchlist/user_service_go/controllers/users"
)

func mapUrls() {
	router.GET("/users/:user_id", users.Get)
	router.GET("/users", users.Find)
	router.POST("/users", users.Insert)
	router.PUT("/users/:user_email", users.Edit)

	router.POST("/login", users.Login)
}
