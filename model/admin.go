package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Admin struct {
	Login     string    `gorm:"not null;size:11;unique" json:"login"`
	Password  string    `gorm:"not null;size:256" json:"password"`
	Name      string    `gorm:"not null;size:50" json:"name"`
	Status    int8      `gorm:"not null;default:1" json:"status"`
	LastLogin time.Time `json:"last_login"`
	gorm.Model
}
