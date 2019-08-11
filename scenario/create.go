package scenario

import (
	"fmt"
	"gorm-example/domain/model"
	"gorm-example/domain/repository"
	"gorm-example/interface/database"

	"github.com/jinzhu/gorm"
)

func Create(db *gorm.DB) error {
	userRepo := database.NewUserRepository(db)

	user := model.User{
		Name:       "Steve Jobs",
		LanguageID: 1,
	}
	if err := add(userRepo, user); err != nil {
		return err
	}

	return nil
}

func add(userRepo repository.UserRepository, user model.User) error {
	fmt.Printf("> Add %v\n", user)
	if err := userRepo.Add(user); err != nil {
		return err
	} else {
		return nil
	}
}
