package mqttservice

import (
	"os"

	"go.ebupt.com/lets/app"
)

func Start() {
	c := make(chan os.Signal, 1)
	mosquittoMQTT()
	// chigoMQTT()
	sig := <-c
	app.LLog.Faltal("MqttService Stopped with signal", sig)
}
