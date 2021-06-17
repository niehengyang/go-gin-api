package model

import (
	"errors"
	"go.ebupt.com/lets/app"
	"go.ebupt.com/lets/server/response"
)

//权限
type Permission struct {
	BaseModel
	Name     string `gorm:"size:100;column(name)" json:"name" form:"name"`
	Parent   string `gorm:"size:10;column(parent)" json:"parent" form:"parent"`
	Status   string `gorm:"size:10;column(status)" json:"status" form:"status"`
	Uid      string `gorm:"size:50;column(uid)" json:"uid" form:"uid"`
	Type     string `gorm:"size:10;column(type)" json:"type" form:"type"`
	Url      string `gorm:"size:500;column(url)" json:"url" form:"url"`
	Icon     string `gorm:"size:500;column(icon)" json:"icon" form:"icon"`
	Describe string `gorm:"size:200;column(describe)" json:"describe" form:"describe"`
}

//权限树
type MenuTree struct {
	Name     string      `json:"name"`
	Parent   string      `json:"parent"`
	Status   string      `json:"status"`
	Uid      string      `json:"uid"`
	Type     string      `json:"type"`
	Url      string      `json:"url"`
	Icon     string      `json:"icon"`
	Describe string      `json:"describe"`
	Children []*MenuTree `json:"children"`
}

func (per Permission) TableName() string {
	return "dd_permission"
}

type SearchPermission struct {
	Name        string
	Page        int
	PageSize    int
	CurrentUser *Admin
}

func (p *SearchPermission) List() *response.PageOption {
	query := app.LDB.Scopes(NameFilter(p.Name), PermissionVisible(p.CurrentUser))

	return &response.PageOption{
		DB:      query,
		Page:    p.Page,
		Limit:   p.PageSize,
		OrderBy: []string{"created_at desc"},
	}
}

func PermissionsTree() ([]interface{}, error) {
	permTree, err := GetMenuTree()
	if err != nil {
		return permTree, errors.New("get permissions error")
	}
	return permTree, nil
}

//构建树形结构
func GetMenuTree() (dataList []interface{}, err error) {
	var parentList []Permission
	//获取父节点

	app.LDB.Where("type = ?", "01").Find(&parentList)

	for _, v := range parentList {
		parent := MenuTree{v.Name, v.Parent, v.Status, v.Uid, v.Type, v.Url, v.Icon, v.Describe, []*MenuTree{}}
		var childrenList []Permission

		app.LDB.Where("parent = ?", v.ID).Find(&childrenList)

		for _, c := range childrenList {
			child := MenuTree{c.Name, c.Parent, c.Status, c.Uid, c.Type, c.Url, c.Icon, c.Describe, []*MenuTree{}}
			parent.Children = append(parent.Children, &child)
		}
		dataList = append(dataList, parent)
	}
	return dataList, nil
}
