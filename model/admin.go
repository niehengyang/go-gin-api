package model

import (
	"errors"
	"go.ebupt.com/lets/app"
)

type Admin struct {
	BaseModel
	Username  string `gorm:"not null;size:11;unique" json:"username"`
	Password  string `gorm:"not null;size:256" json:"_"`
	Name      string `gorm:"not null;size:50" json:"name"`
	Status    int8   `gorm:"not null;default:1" json:"status"`
	Type      string `gorm:"size:20" json:"type"`
	PhoneNum  string `gorm:"not null;unique;size:50" json:"phone_num"`
	Email     string `gorm:"size:50" json:"email"`
	Avatar    string `gorm:"size:256" json:"avatar"`
	Describe  string `gorm:"size:200" json:"describe"`
	Token     string `gorm:"size:256" json:"_"`
	TokenEx   string `gorm:"size:256" json:"_"`
	LastLogin XTime  `json:"last_login"`

	Roles []Role `gorm:"many2many:user_roles"`
}

func (a Admin) TableName() string {
	return "dd_admin"
}

/*
	账号类别，0：系统账号，1：厂区账号
*/
const Type_SystemAccount = "0"
const Type_FactoryAccount = "1"

/*
	角色，0：操作管理员（具有操作与查看权限），1：视图管理员（只有查看权限）
*/
const Role_ReadAndWrite = 0
const Role_ReadOnly = 1

/*--------- help methods -----------*/
/*
	是否系统账号
*/
func IsSystemAccount(account *Admin) bool {
	if account.Type == Type_SystemAccount {
		return true
	} else {
		return false
	}
}

func (a *Admin) Create() (uint, error) {

	result := app.LDB.Create(&a)
	if result.Error != nil {
		return 0, errors.New("user create fail")
	}

	id := a.BaseModel.ID

	//更新关系
	app.LDB.Model(&a).Association("Roles").Append(a.Roles)

	return id, nil
}
