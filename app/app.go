package app

import (
	"github.com/gin-gonic/gin"
	"github.com/muchlist/user_service_go/db"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp memulai http server
func StartApp() {

	client, ctx, cancel := db.Init()
	defer client.Disconnect(ctx)
	defer cancel()

	// mapping urls file url_mappings.go
	mapUrls()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
