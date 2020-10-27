package app

import (
	"github.com/muchlist/user_service_go/controllers/middleware"
	"github.com/muchlist/user_service_go/controllers/ping_controller"
	"github.com/muchlist/user_service_go/controllers/user_controller"
)

func mapUrls() {

	router.Static("/images", "./static/images")

	api := router.Group("/api/v1")

	api.POST("/login", user_controller.Login)
	api.GET("/users/:user_id", middleware.AuthMiddleware, user_controller.Get)
	api.GET("/users", middleware.AuthMiddleware, user_controller.Find)
	api.POST("/users", middleware.AuthMiddleware, user_controller.Insert)
	api.PUT("/users/:user_email", middleware.AuthMiddleware, user_controller.Edit)
	api.POST("/avatar", middleware.AuthMiddleware, user_controller.UploadImage)

	api.GET("/ping", ping_controller.Ping)

}
