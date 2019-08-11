package scenario

import (
	"fmt"
	"gorm-example/interface/database"

	"github.com/jinzhu/gorm"
)

func Delete(db *gorm.DB) error {
	userRepo := database.NewUserRepository(db)

	user, err := userRepo.Last()
	if err != nil {
		return err
	}
	printRecord(user)

	id := user.ID
	fmt.Printf("> Delete (%v)\n", id)
	if err := userRepo.Delete(id); err != nil {
		return err
	}

	user, err = userRepo.Last()
	if err != nil {
		return err
	}
	printRecord(user)

	return nil
}
