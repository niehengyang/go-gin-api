package v1

import (
	"github.com/gin-gonic/gin"
	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/server/response"
	"gotutorial/model"
	"gotutorial/utils"
)

func PermissionList(c *gin.Context) {
	R := response.New(c)
	var permissions []model.Permission

	name := c.DefaultQuery("name", "")
	page := utils.StringToInt(c.DefaultQuery("page", "1"))
	pageSize := utils.StringToInt(c.DefaultQuery("pageSize", app.AppConfig.PageSize))
	currentUser := GetLoginAccount(c)

	params := model.SearchPermission{Name: name, Page: page, PageSize: pageSize, CurrentUser: currentUser}

	pageOption := params.List()

	R.Pagination(pageOption, &permissions, R.AfterPaginator)
}

func PermissionsTree(c *gin.Context) {
	R := response.New(c)
	permissions, err := model.PermissionsTree()
	if err != nil {
		R.Error500("查询失败")
		return
	}

	R.Data(permissions)
}
