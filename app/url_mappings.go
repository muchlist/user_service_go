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
	api.GET("/ping", ping_controller.Ping)

	apiAuth := router.Group("/api/v1")
	apiAuth.Use(middleware.AuthMiddleware)
	apiAuth.GET("/users/:user_id", user_controller.Get)
	apiAuth.GET("/profile", user_controller.GetProfile)
	apiAuth.GET("/users", user_controller.Find)
	apiAuth.POST("/users", user_controller.Insert)
	apiAuth.POST("/avatar", user_controller.UploadImage)
	apiAuth.POST("/profile/change-password", user_controller.ChangePassword)

	apiAuthAdmin := router.Group("/api/v1/admin")
	apiAuthAdmin.Use(middleware.AuthAdminMiddleware)
	apiAuthAdmin.PUT("/users/:user_email", user_controller.Edit)
	apiAuthAdmin.DELETE("/users/:user_email", user_controller.Delete)
	apiAuthAdmin.GET("/users/:user_email/reset-password", user_controller.ResetPassword)

}
