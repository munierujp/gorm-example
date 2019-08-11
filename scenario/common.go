package scenario

import (
	"fmt"
	"gorm-example/domain/model"
	"gorm-example/domain/repository"
)

const (
	header = "id" + "\t" + "name" + "\t" + "language_id" + "\t" + "created_at" + "\t" + "updated_at" + "\t" + "deleted_at"
)

func printLastReocrd(userRepo repository.UserRepository) error {
	user, err := userRepo.Last()
	if err != nil {
		return err
	}
	printRecord(user)
	return nil
}

func printRecord(user *model.User) {
	fmt.Println(header)
	row := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", user.ID, user.Name, user.LanguageID, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	fmt.Println(row)
}
