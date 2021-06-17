package v1

import (
	"github.com/gin-gonic/gin"
	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/server/request"
	"go.ebupt.com/lets/server/response"
	"gotutorial/model"
	"gotutorial/utils"
)

func RoleList(c *gin.Context) {
	R := response.New(c)
	var params model.SearchRole
	var roles []model.Role

	params.Name = c.DefaultQuery("name", "")
	params.Page = utils.StringToInt(c.DefaultQuery("page", "1"))
	params.PageSize = utils.StringToInt(c.DefaultQuery("pageSize", app.AppConfig.PageSize))
	params.CurrentUser = GetLoginAccount(c)

	pageOption := params.List()

	R.Pagination(pageOption, &roles, R.AfterPaginator)
}

func CreateRole(c *gin.Context) {
	R := response.New(c)

	var form model.Role
	success, errors := request.BindAndValid(c, &form)

	if !success {
		R.Error400(errors)
		return
	}

	loginUser := GetLoginAccount(c)
	form.UserId = loginUser.BaseModel.ID

	//创建
	_, result := form.Create()

	if result != nil {
		R.Error500(result)
		return
	}

	R.Success("创建成功")
}

func EditRole(c *gin.Context) {
	var form model.Role
	R := response.New(c)

	roleId := c.Param("id")
	success, errors := request.BindAndValid(c, &form)

	if !success {
		R.Error400(errors)
		return
	}

	//编辑
	result := form.Edit(utils.StringToInt(roleId))
	if result != nil {
		R.Error500(result)
		return
	}
	R.Success("edit success")
}

func RoleItem(c *gin.Context) {
	var form model.Role
	R := response.New(c)

	id := c.Param("id")

	//查找
	item, result := form.Item(utils.StringToInt(id))

	if result != nil {
		R.Error500(result)
		return
	}
	R.Data(item)
}

func DeleteRole(c *gin.Context) {
	var form model.Role
	R := response.New(c)

	roleId := c.Param("id")
	err := form.Delete(utils.StringToInt(roleId))

	if err != nil {
		R.Error500("删除失败")
		return
	}
	R.Success("删除成功")
}

func BindPermissions(c *gin.Context) {
	var permissionUids []string
	R := response.New(c)
	roleId := c.Param("id")

	err := c.ShouldBindJSON(&permissionUids)
	if err != nil {
		R.Error400("参数错误")
		return
	}

	result := model.BindPermissions(utils.StringToInt(roleId), permissionUids)

	if result != nil {
		R.Error500(result)
		return
	}
	R.Success("绑定成功")
}

func GetRolePermissions(c *gin.Context) {
	id := c.Param("id")
	R := response.New(c)

	//查找
	permissions, err := model.GetRolePermissions(utils.StringToInt(id))

	if err != nil {
		R.Error500("查询失败")
		return
	}
	R.Data(permissions)
}
