package database

import (
	"gorm-example/domain/model"
	"gorm-example/domain/repository"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
