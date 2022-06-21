package model

import (
	"time"
)

// gin_vuePQ/model/user.go

// User 获取帖子列表query string参数
type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `form:"name" gorm:"type:varchar(20);not null"`
	Telephone string `form:"telephone" gorm:"varchar(11);not null;unique"`
	Password  string `form:"password" gorm:"size:255;not null"`
}
