package repository

import (
	"gorm-example/domain/model"
)

type UserRepository interface {
	FindByID(uint) (*model.User, error)
}
