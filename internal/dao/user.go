package dao

import "gorm.io/gorm"

type UserDao struct {
	db *gorm.DB
}