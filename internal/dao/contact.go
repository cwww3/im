package dao

import "gorm.io/gorm"

type ContactDao struct {
	db *gorm.DB
}