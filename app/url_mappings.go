package app

import (
	"github.com/muchlist/user_service_go/controllers/middleware"
	"github.com/muchlist/user_service_go/controllers/user_handler"
)

func mapUrls() {

	api := router.Group("/api/v1")

	api.GET("/users/:user_id", user_handler.Get)
	api.GET("/users", middleware.AuthMiddleware, user_handler.Find)
	api.POST("/users", user_handler.Insert)
	api.PUT("/users/:user_email", user_handler.Edit)
	api.POST("/login", user_handler.Login)

}
