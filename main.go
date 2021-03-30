package main

import (
	"gotutorial/model"
	"gotutorial/router"

	"go.ebupt.com/lets/server"
)

func main() {
	apiServer := server.New("./config/config.toml")
	server.LDB.AutoMigrate(&model.Admin{})
	apiServer.BindRouter(router.RouterBinder)
	apiServer.Go()
}
