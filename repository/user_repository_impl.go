package repository

import (
	"basecode/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (u *UserRepositoryImpl) Save(user model.User) error {
	err := u.DB.Create(&user).Error

	return err
}

func (u *UserRepositoryImpl) FindOneByUsername(username string) (model.User, error) {
	var user model.User

	err := u.DB.Where("username = ? ", username).First(&user).Error

	return user, err
}

func (u *UserRepositoryImpl) UpdateToken(user model.User) error {
	err := u.DB.Save(&user).Error

	return err
}
