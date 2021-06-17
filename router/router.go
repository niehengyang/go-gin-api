package router

import (
	v1 "gotutorial/app/api/v1"
	"gotutorial/middleware"

	"github.com/gin-gonic/gin"
)

func RouterBinder(ge *gin.Engine) {

	ge.GET("/", v1.ServerAbout)

	//登录
	ge.POST("/auth/login", v1.Login)

	ge.Use(middleware.JwtAuth())
	ge.GET("/admin", v1.GetAdmin)
	ge.POST("/admin", v1.CreateAdmin)

	//角色
	ge.POST("/role", v1.CreateRole)
	ge.PUT("/role/:id", v1.EditRole)
	ge.GET("/role/:id", v1.RoleItem)
	ge.GET("/role", v1.RoleList)
	ge.DELETE("/role/:id", v1.DeleteRole)
	ge.PUT("/bindpermissions/:id", v1.BindPermissions)
	ge.GET("/rolepermissions/:id", v1.GetRolePermissions)

	//权限
	ge.GET("/permissionlist", v1.PermissionList)
	ge.GET("/permissionstree", v1.PermissionsTree)
}
