package main

import (
	"fmt"
	"gotutorial/app/mqttservice"
	"gotutorial/router"

	"github.com/spf13/cobra"
	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/server"
)

var rootCmd = &cobra.Command{
	Use:   "lets",
	Short: "lets 业务框架命令",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("请指定要运行的子命令")
	},
}

var mqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "运行MQTT服务",
	Run: func(cmd *cobra.Command, args []string) {
		startMqttClient()
	},
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "运行Api服务",
	Run: func(cmd *cobra.Command, args []string) {
		StartApiServer()
	},
}

func main() {

	app.Bootstrap("./config/config.toml")

	rootCmd.AddCommand(mqttCmd)
	rootCmd.AddCommand(apiCmd)
	rootCmd.Execute()
}

func StartApiServer() {
	apiServer := server.New()
	// server.LDB.AutoMigrate(&model.Admin{})
	apiServer.BindRouter(router.RouterBinder)
	apiServer.Go()
}

func startMqttClient() {
	mqttservice.Start()
}

func StartRabbitMQConsumer() {

}
