package scenario

import (
	"fmt"
	"gorm-example/interface/database"

	"github.com/jinzhu/gorm"
)

func Update(db *gorm.DB) error {
	userRepo := database.NewUserRepository(db)

	user, err := userRepo.Last()
	if err != nil {
		return err
	}
	printRecord(user)

	user.Name = user.Name + "😄"
	fmt.Printf("> Update %v\n", user)
	if err := userRepo.Update(user); err != nil {
		return err
	}

	user, err = userRepo.Last()
	if err != nil {
		return err
	}
	printRecord(user)

	return nil
}
