package middleware

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.ebupt.com/lets/app"
	"gotutorial/model"

	"github.com/gin-gonic/gin"
	"go.ebupt.com/lets/server/auth"
	"go.ebupt.com/lets/server/helper"
	"go.ebupt.com/lets/server/response"
)

func JwtAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		R := response.New(c)

		token := c.GetHeader("Authorization")
		if token == "" {
			R.Error401(response.ResponseMap{
				"code":    auth.ExpireOrErrorToken,
				"message": auth.AuthErrorMap[auth.ExpireOrErrorToken],
			})
			c.Abort()
			return
		}

		claims, err := helper.JwtParseToken(token)
		if err != nil {
			R.Error500(fmt.Sprintf("解析Token失败%v", err))
			c.Abort()
			return
		}

		setLoginUser(c, token)

		fmt.Println(claims.Identify)
		c.Next()
	}
}

func setLoginUser(c *gin.Context, token string) {
	var admin model.Admin
	R := response.New(c)
	accountRes := app.LDB.Where("token = ?", token).First(&admin)
	if accountRes.Error != nil && errors.Is(accountRes.Error, gorm.ErrRecordNotFound) {
		R.Error401(response.ResponseMap{
			"code":    auth.ExpireOrErrorToken,
			"message": auth.AuthErrorMap[auth.ExpireOrErrorToken],
		})
		c.Abort()
		return
	}
	c.Set("loginUser", &admin)

	return
}
