package command

import (
	"github.com/spf13/cobra"
	"go.ebupt.com/lets/server"
	"gotutorial/router"
)

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "运行Api服务",
	Run: func(cmd *cobra.Command, args []string) {
		StartApiServer()
	},
}

func StartApiServer() {
	apiServer := server.New()
	apiServer.BindRouter(router.RouterBinder)
	apiServer.Go()
}
