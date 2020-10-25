package app

import (
	"github.com/muchlist/user_service_go/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
	router.POST("/users", controllers.InsertUser)
}
