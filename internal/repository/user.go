package repository

import (
	"github.com/cwww3/im/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindByEmail(email string) (model.User,error) {
	var user model.User
	err := ur.db.Model(new(model.User)).Where("email = ?",email).First(&user).Error
	return user,err
}

func (ur *UserRepository) FindAll() ([]model.User,error) {
	var users []model.User
	err := ur.db.Model(new(model.User)).Find(&users).Error
	return users,err
}

func (ur *UserRepository) FindMessageContactsByOwnerUidOrderByMidDesc(userId uint64) ([]model.User,error) {
	var users []model.User
	err := ur.db.Model(new(model.MessageContact)).Where("owner_uid = ?",userId).Order("mid desc").Find(&users).Error
	return users,err
}





