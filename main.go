package main

import (
	"fmt"
	"gotutorial/command"

	"github.com/spf13/cobra"
	"go.ebupt.com/lets/app"
)

var rootCmd = &cobra.Command{
	Use:   "lets",
	Short: "lets 业务框架命令",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("请指定要运行的子命令")
	},
}

func main() {

	app.Bootstrap("./config/config.toml")

	rootCmd.AddCommand(command.MqttCmd)
	rootCmd.AddCommand(command.ApiCmd)
	rootCmd.AddCommand(command.MigrateCmd)

	rootCmd.Execute()
}
