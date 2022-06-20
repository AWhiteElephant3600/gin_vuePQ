package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `form:"name" gorm:"type:varchar(20);not null"`
	Telephone string `form:"telephone" gorm:"varchar(11);not null;unique"`
	Password string `form:"password" gorm:"size:255;not null"`
}