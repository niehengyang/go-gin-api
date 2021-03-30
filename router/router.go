package router

import (
	v1 "gotutorial/app/api/v1"
	"gotutorial/middleware"

	"github.com/gin-gonic/gin"
)

func RouterBinder(ge *gin.Engine) {

	ge.POST("/auth/login", v1.Login)

	ge.Use(middleware.JwtAuth())
	ge.GET("/admin", v1.GetAdmin)
	ge.POST("/admin", v1.CreateAdmin)

}
