package app

import (
	"github.com/muchlist/user_service_go/controllers/middleware"
	"github.com/muchlist/user_service_go/controllers/user_handler"
)

func mapUrls() {

	api := router.Group("/api/v1")

	api.POST("/login", user_handler.Login)
	api.GET("/users/:user_id", middleware.AuthMiddleware, user_handler.Get)
	api.GET("/users", middleware.AuthMiddleware, user_handler.Find)
	api.POST("/users", middleware.AuthMiddleware, user_handler.Insert)
	api.PUT("/users/:user_email", middleware.AuthMiddleware, user_handler.Edit)
	api.POST("/avatar", middleware.AuthMiddleware, user_handler.UploadImage)

}
