package scenario

import (
	"fmt"
	"gorm-example/interface/database"

	"github.com/jinzhu/gorm"
)

func Read(db *gorm.DB) error {
	userRepo := database.NewUserRepository(db)

	user, err := userRepo.Last()
	if err != nil {
		return err
	}

	id := user.ID
	fmt.Printf("> Find by id (%v)\n", id)
	user, err = userRepo.FindByID(id)
	if err != nil {
		return err
	}
	printRecord(user)

	return nil
}
