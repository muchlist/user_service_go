package main

import (
	"github.com/muchlist/user_service_go/app"
	"github.com/muchlist/user_service_go/db"
)

func main() {
	db.Init()
	app.StartApp()
}
