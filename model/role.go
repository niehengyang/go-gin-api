package model

import (
	"errors"
	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/server/response"
)

//查询条件
type SearchRole struct {
	Name        string
	Page        int
	PageSize    int
	CurrentUser *Admin
}

//角色
type Role struct {
	BaseModel
	Name       string       `gorm:"not null;size:100;column(name)" valid:"Required" json:"name" form:"name"`
	UserId     uint         `gorm:"not null" json:"user_id" form:"user_id"`
	Desc       string       `gorm:"size:500;column(desc)" json:"desc" form:"desc"`
	Permission []Permission `gorm:"many2many:role_permissions"`
}

func (role Role) TableName() string {
	return "dd_role"
}

func (r *Role) Create() (uint, error) {
	result := app.LDB.Create(&r)
	id := r.BaseModel.ID
	if result.Error != nil {
		return 0, result.Error
	}

	return id, nil
}

func (r *Role) Edit(roleId int) error {

	var role Role
	app.LDB.First(&role, roleId)

	result := app.LDB.Model(&role).Update(r)
	if result.Error != nil {
		return errors.New("edit fail")
	}

	return nil
}

func (r *Role) Item(roleId int) (*Role, error) {
	result := app.LDB.Preload("Permission").First(&r, roleId)
	if result.Error != nil {
		return r, errors.New("查询失败")
	}
	return r, nil
}

func (s *SearchRole) List() *response.PageOption {
	query := app.LDB.Scopes(NameFilter(s.Name), DataVisible(s.CurrentUser))

	return &response.PageOption{
		DB:      query,
		Page:    s.Page,
		Limit:   s.PageSize,
		OrderBy: []string{"created_at desc"},
	}
}

func (r *Role) Delete(id int) error {
	result := app.LDB.First(&r, id)
	if err := result.Error; err != nil {
		return errors.New("role not find")
	}

	//清除关联
	app.LDB.Model(&r).Association("Permission").Clear()

	//彻底删除
	app.LDB.Unscoped().Delete(&r)

	return nil
}

func BindPermissions(id int, perUids []string) error {
	var role Role
	var permissions []Permission

	app.LDB.Where("uid in (?)", perUids).Find(&permissions)

	result := app.LDB.First(&role, id)
	if err := result.Error; err != nil {
		return errors.New("role not find")
	}

	app.LDB.Model(&role).Association("Permission").Replace(permissions)

	return nil
}

func GetRolePermissions(id int) ([]string, error) {
	var role Role
	var uids []string

	result := app.LDB.First(&role, id)
	if err := result.Error; err != nil {
		return uids, errors.New("role not find")
	}

	var permissions []Permission
	app.LDB.Model(&role).Association("Permission").Find(&permissions)

	for _, v := range permissions {
		uids = append(uids, v.Uid)
	}

	return uids, nil
}
