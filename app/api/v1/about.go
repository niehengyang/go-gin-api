package v1

import (
	"github.com/gin-gonic/gin"
	"go.ebupt.com/lets/server/response"
)

func ServerAbout(ctx *gin.Context) {

	R := response.New(ctx)
	R.Success("Lets Project is running")

}
