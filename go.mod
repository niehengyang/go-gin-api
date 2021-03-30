module gotutorial

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	go.ebupt.com/lets v0.0.1
)

replace go.ebupt.com/lets => ../../../go/src/go.ebupt.com/lets/
