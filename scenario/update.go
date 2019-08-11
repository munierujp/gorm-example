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

	name := user.Name + "😄"
	user.Name = name
	fmt.Printf("> Update name to \"%v\"\n", name)
	if err := userRepo.Update(user); err != nil {
		return err
	}
	printLastReocrd(userRepo)

	return nil
}
