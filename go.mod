module gotutorial

go 1.15

require (
	github.com/eclipse/paho.mqtt.golang v1.3.2
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	github.com/spf13/cobra v1.1.3
	go.ebupt.com/lets v0.0.1
)

replace go.ebupt.com/lets => ../../go/src/lets/
