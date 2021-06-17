package v1

import (
	"fmt"
	"gotutorial/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/server/auth"
	"go.ebupt.com/lets/server/helper"
	"go.ebupt.com/lets/server/request"
	"go.ebupt.com/lets/server/response"
)

type LoginForm struct {
	Username string `valid:"Required" form:"username"`
	Password string `valid:"Required" form:"password"`
}

func Login(c *gin.Context) {

	app.LRedis.Set(app.LRedisCtx, "login_key", "Hello world", 0)

	R := response.New(c)
	var form LoginForm
	admin := new(model.Admin)
	success, errors := request.BindAndValid(c, &form)

	fmt.Println(success)

	if !success {
		R.Error400(errors)
		return
	}

	result := app.LDB.Where("username = ?", form.Username).First(&admin)

	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			R.Error401(response.ResponseMap{
				"code":    auth.ErrorLoginOrPassword,
				"message": auth.AuthErrorMap[auth.ErrorLoginOrPassword],
			})
			return
		} else {
			R.Error500(result.Error)
			return
		}
	}

	if admin.Status != 1 {
		R.Error401(response.ResponseMap{
			"code":    auth.UserIsDisabled,
			"message": auth.AuthErrorMap[auth.UserIsDisabled],
		})
		return
	}

	err := auth.ComparePassword(admin.Password, form.Password)
	if err != nil {
		R.Error401(response.ResponseMap{
			"code":    auth.ErrorLoginOrPassword,
			"message": auth.AuthErrorMap[auth.ErrorLoginOrPassword],
		})
		return
	}

	jwtIdentify := map[string]interface{}{
		"user_id":        admin.BaseModel.ID,
		"other_identify": "otheridenfity",
	}

	jwtToken, err := helper.JwtGenerateToken(jwtIdentify, app.AppConfig.JwtIssuer, 3*time.Hour)

	if err != nil {
		R.Error500(fmt.Sprintf("生成JwtToken失败%v", err))
		return
	}

	admin.Token = jwtToken
	admin.LastLogin = model.XTime{time.Now()}
	app.LDB.Save(&admin)

	R.Success(response.ResponseMap{
		"code":    200,
		"token":   jwtToken,
		"message": "登录成功",
	})

}
func Logout(c *gin.Context) {

}

/**
获取当前登录账号
*/
func GetLoginAccount(c *gin.Context) *model.Admin {
	accountInterface := c.MustGet("loginUser")
	loginAccount := accountInterface.(*model.Admin)
	return loginAccount
}
