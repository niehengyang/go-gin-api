package main

import (
	"gotutorial/model"
	"gotutorial/router"

	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/server"
)

func main() {

	app.Bootstrap("./config/config.toml")

	apiServer := server.New()
	app.LDB.AutoMigrate(&model.Admin{})
	apiServer.BindRouter(router.RouterBinder)
	apiServer.Go()
}
