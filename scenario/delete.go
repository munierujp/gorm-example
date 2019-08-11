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

	id := user.ID
	fmt.Printf("> Delete (%v)\n", id)
	if err := userRepo.Delete(id); err != nil {
		return err
	}
	printLastReocrd(userRepo)

	return nil
}
