package scenario

import (
	"fmt"
	"gorm-example/interface/database"
	"gorm-example/domain/repository"

	"github.com/jinzhu/gorm"
)

func Read(db *gorm.DB) error {
	userRepo := database.NewUserRepository(db)

	if err := findByID(userRepo, 1); err != nil {
		return err
	}

	if err := findByID(userRepo, 3); err != nil {
		return err
	}

	return nil
}

func findByID(userRepo repository.UserRepository, id uint) error {
	fmt.Printf("> Find by id (%v)\n", id)
	if user, err := userRepo.FindByID(id); err != nil {
		return err
	} else {
		fmt.Println("id" + "\t" + "name" + "\t" + "language_id" + "\t" + "created_at" + "\t" + "updated_at" + "\t" + "deleted_at")
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", user.ID, user.Name, user.LanguageID, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
		return nil
	}
}
