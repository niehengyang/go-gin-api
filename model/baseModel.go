package model

import (
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// 1. 创建 time.Time 类型的副本 XTime；
type XTime struct {
	time.Time
}

const TimeFormat = "2006-01-02 15:04:05"

func (t *XTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = XTime{now}
	return err
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(output), nil
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 5. 自定义 BaseModel，结构和 gorm.Model 一致，将 time.Time 替换为 LocalTime；
type BaseModel struct {
	ID        uint   `gorm:"primary_key,AUTO_INCREMENT" json:"id"`
	CreatedAt XTime  `json:"created_at"`
	UpdatedAt XTime  `json:"updated_at"`
	DeletedAt *XTime `sql:"index" json:"deleted_at"`
}

/**
根据名称过滤
*/
func NameFilter(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name LIKE ?", "%"+name+"%")
	}
}

/**
根据状态过滤
*/
func StatusFilter(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status == "" {
			return db
		}
		return db.Where("status = ?", status)
	}
}

/*
数据可见性过滤
*/
func DataVisible(account *Admin) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !IsSystemAccount(account) {
			//TODO: 限制条件
			return db
		}
		return db
	}
}

/*
 权限过滤
*/
func PermissionVisible(account *Admin) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !IsSystemAccount(account) {
			//TODO: 限制条件
			return db
		}
		return db
	}
}
