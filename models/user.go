package models

import "time"

type User struct {
	ID       uint       `gorm:"colume:id;PRIMERY_KEY"`
	UserName string     `gorm:"colume:username"`
	Password string     `gorm:"colume:password"`
	FullName string     `gorm:"colume:fullname"`
	CreateAt *time.Time `grom:"autoCreateTime:milli"`
	UpdateAt *time.Time `gorm:"autoUpdateTime:milli"`
}
