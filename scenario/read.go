package scenario

import (
	"fmt"
	"gorm-example/interface/database"

	"github.com/jinzhu/gorm"
)

func Read(db *gorm.DB) error {
	userRepo := database.NewUserRepository(db)

	fmt.Println("> Get last record")
	user, err := userRepo.Last()
	if err != nil {
		return err
	}
	fmt.Println("id" + "\t" + "name" + "\t" + "language_id" + "\t" + "created_at" + "\t" + "updated_at" + "\t" + "deleted_at")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", user.ID, user.Name, user.LanguageID, user.CreatedAt, user.UpdatedAt, user.DeletedAt)

	id := user.ID
	fmt.Printf("> Find by id (%v)\n", id)
	user, err = userRepo.FindByID(id)
	if err != nil {
		return err
	}
	fmt.Println("id" + "\t" + "name" + "\t" + "language_id" + "\t" + "created_at" + "\t" + "updated_at" + "\t" + "deleted_at")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", user.ID, user.Name, user.LanguageID, user.CreatedAt, user.UpdatedAt, user.DeletedAt)

	return nil
}
