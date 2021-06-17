package Seeds

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/server/auth"
	"gotutorial/model"
	"time"
)

func InitSeed() {
	seedAdmin()       //生成超管
	seedPermissions() //生成权限
}

func seedAdmin() {
	/*---------- Seed 初始管理员数据 --------------*/
	secretPassword, err := auth.HashPassword("eb823258")
	if err != nil {
		fmt.Sprintf("密码加密错误:%v", err)
		return
	}

	initAdmins := []model.Admin{
		{Username: "superadmin", Password: string(secretPassword), Name: "超级管理员", PhoneNum: "13999999999"},
	}
	for _, admin := range initAdmins {
		account := model.Admin{}
		adminQuery := app.LDB.Where("username = ?", admin.Username).First(&account)

		if adminQuery.Error != nil && errors.Is(adminQuery.Error, gorm.ErrRecordNotFound) {

			seedAdmin := model.Admin{Username: admin.Username, Password: string(secretPassword), Name: admin.Name, Type: "0", PhoneNum: admin.PhoneNum, LastLogin: model.XTime{time.Now()}}
			result := app.LDB.Create(&seedAdmin)
			if result.Error != nil {
				fmt.Println("数据库插入错误>", result.Error)
				return
			}
		}
	}
}

func seedPermissions() {

	var perArr [9]model.Permission

	perArr[0] = model.Permission{Name: "首页", Parent: "0", Status: "1", Type: "01", Uid: "home", Url: "test", Icon: "test", Describe: "首页"}
	perArr[1] = model.Permission{Name: "系统管理", Parent: "0", Status: "1", Type: "01", Uid: "system", Url: "test", Icon: "test", Describe: "系统管理"}
	perArr[2] = model.Permission{Name: "用户管理", Parent: "2", Status: "1", Type: "02", Uid: "user", Url: "test", Icon: "test", Describe: "用户管理"}
	perArr[3] = model.Permission{Name: "日志管理", Parent: "2", Status: "1", Type: "02", Uid: "log", Url: "test", Icon: "test", Describe: "日志管理"}
	perArr[4] = model.Permission{Name: "模型管理", Parent: "0", Status: "1", Type: "01", Uid: "model", Url: "test", Icon: "test", Describe: "模型管理"}
	perArr[5] = model.Permission{Name: "应用管理", Parent: "0", Status: "1", Type: "01", Uid: "app", Url: "test", Icon: "test", Describe: "应用管理"}
	perArr[6] = model.Permission{Name: "角色管理", Parent: "2", Status: "1", Type: "02", Uid: "role", Url: "test", Icon: "test", Describe: "角色管理"}
	perArr[7] = model.Permission{Name: "分组管理", Parent: "5", Status: "1", Type: "02", Uid: "model_group", Url: "test", Icon: "test", Describe: "分组管理"}
	perArr[8] = model.Permission{Name: "模型列表", Parent: "5", Status: "1", Type: "02", Uid: "model_list", Url: "test", Icon: "test", Describe: "模型列表"}

	perArr[0].BaseModel.ID = 1
	perArr[1].BaseModel.ID = 2
	perArr[2].BaseModel.ID = 3
	perArr[3].BaseModel.ID = 4
	perArr[4].BaseModel.ID = 5
	perArr[5].BaseModel.ID = 6
	perArr[6].BaseModel.ID = 7
	perArr[7].BaseModel.ID = 8
	perArr[8].BaseModel.ID = 9

	for i := 0; i < len(perArr); i++ {
		app.LDB.FirstOrCreate(&perArr[i])
	}
}
