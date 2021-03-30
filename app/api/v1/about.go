package v1

import (
	"gotutorial/app/mqttservice"

	"github.com/gin-gonic/gin"
	"go.ebupt.com/lets/server/response"
)

func ServerAbout(ctx *gin.Context) {

	mqttservice.MosquittoPub("topic/test", 1, false, "from api server")

	R := response.New(ctx)
	R.Success("Lets Project is running")

}
