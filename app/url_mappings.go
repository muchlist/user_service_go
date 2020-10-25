package app

import (
	"github.com/muchlist/user_service_go/controllers/users"
)

func mapUrls() {
	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Insert)
}
