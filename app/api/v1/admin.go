package v1

import (
	"fmt"
	"gotutorial/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.ebupt.com/lets/server"
	"go.ebupt.com/lets/server/auth"
	"go.ebupt.com/lets/server/request"
	"go.ebupt.com/lets/server/response"
)

func GetAdmin(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", server.AppConfig.PageSize))

	var admins []model.Admin

	r := response.New(c)

	r.Pagination(&response.PageOption{
		DB:      server.LDB,
		Page:    page,
		Limit:   pageSize,
		OrderBy: []string{"created_at desc"},
	}, &admins)

}

type AdminCreateForm struct {
	Name     string `valid:"Required" form:"name"`
	Status   int8   `valid:"Range(0,1)" form:"status"`
	Login    string `valid:"Required;Mobile" form:"login"`
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
	admin.Login = form.Login
	secretPassword, err := auth.HashPassword(form.Password)
	if err != nil {
		R.Error500(fmt.Sprintf("密码加密错误:%v", err))
		return
	}
	admin.Password = string(secretPassword)
	admin.Status = form.Status
	admin.LastLogin = time.Now()

	result := server.LDB.Create(&admin)

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
