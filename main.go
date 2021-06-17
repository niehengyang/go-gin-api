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
		fmt.Printf("---------------指令集-----------------\n")
		fmt.Printf("|          api  -> 接口服务           |\n")
		fmt.Printf("--------------------------------------\n")
		fmt.Printf("|          rbmq -> 队列服务           |\n")
		fmt.Printf("--------------------------------------\n")
		fmt.Printf("|          mqtt -> 数据服务           |\n")
		fmt.Printf("--------------------------------------\n")
		fmt.Printf("请指定要运行的子命令:\n")
	},
}

func main() {

	app.Bootstrap("./config/config.toml")

	rootCmd.AddCommand(command.MqttCmd)
	rootCmd.AddCommand(command.ApiCmd)
	rootCmd.AddCommand(command.MigrateCmd)

	rootCmd.Execute()
}
