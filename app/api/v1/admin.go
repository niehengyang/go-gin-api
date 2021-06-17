package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/server/auth"
	"go.ebupt.com/lets/server/request"
	"go.ebupt.com/lets/server/response"
	"gotutorial/model"
	"strconv"
	"time"
)

func GetAdmin(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", app.AppConfig.PageSize))

	var admins []model.Admin

	r := response.New(c)

	r.Pagination(&response.PageOption{
		DB:      app.LDB,
		Page:    page,
		Limit:   pageSize,
		OrderBy: []string{"created_at desc"},
	}, &admins, r.AfterPaginator)
}

type AdminCreateForm struct {
	Name     string `valid:"Required" form:"name"`
	Status   int8   `valid:"Range(0,1)" form:"status"`
	Username string `valid:"Required;Mobile" form:"username"`
	Password string `valid:"Required" form:"password"`
}

func CreateAdmin(c *gin.Context) {

	R := response.New(c)
	var form AdminCreateForm
	success, errors := request.BindAndValid(c, &form)

	if !success {
		R.Error400(errors)
		return
	}

	admin := new(model.Admin)
	admin.Name = form.Name
	admin.Username = form.Username
	secretPassword, err := auth.HashPassword(form.Password)
	if err != nil {
		R.Error500(fmt.Sprintf("密码加密错误:%v", err))
		return
	}
	admin.Password = string(secretPassword)
	admin.Status = form.Status
	admin.LastLogin = model.XTime{time.Now()}

	result := app.LDB.Create(&admin)

	if result.Error != nil {
		R.Error500(result.Error)
		return
	}

	R.Created(response.ResponseMap{
		"id":      admin.ID,
		"message": "Created",
	})

}

func UpdateUser(c *gin.Engine) {

}

func DeleteUser(c *gin.Engine) {

}
