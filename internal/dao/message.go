package dao

import "gorm.io/gorm"

type MessageDao struct {
	db *gorm.DB
}
