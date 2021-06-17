package command

import (
	"github.com/spf13/cobra"
	"gotutorial/app/mqttservice"
)

var MqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "运行MQTT服务",
	Run: func(cmd *cobra.Command, args []string) {
		startMqttClient()
	},
}

func startMqttClient() {
	mqttservice.Start()
}
