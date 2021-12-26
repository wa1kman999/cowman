package model

import (
	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null" json:"username" binding:"required,max=12" label:"用户名"`
	Password  string `gorm:"type:varchar(500);not null" json:"-" binding:"required,min=6,max=120" label:"密码"`
	Role      int    `gorm:"type:int;DEFAULT:2" json:"role" label:"角色码"`
	Phone     string `gorm:"type:varchar(11);null" label:"电话"`
	Email     string `gorm:"type:varchar(20);null" label:"邮箱"`
	HeaderImg string `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Boss      string `gorm:"type:varchar(20);not null" json:"boss"`
}
