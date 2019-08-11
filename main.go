package main

import (
	"fmt"
	"gorm-example/scenario"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USER_NAME = "user"
	PASSWORD  = "pass"
	DATABASE  = "sample"
	HOST      = "localhost"
	PORT      = "3306"
	CHARSET   = "utf8mb4"
	COLLATION = "utf8mb4_bin"
)

func main() {
	db, err := gorm.Open("mysql", USER_NAME+":"+PASSWORD+"@tcp("+HOST+":"+PORT+")/"+DATABASE+"?charset="+CHARSET+"&collation="+COLLATION+"&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create
	fmt.Println("> Create")
	if err := scenario.Create(db); err != nil {
		panic(err)
	}

	// Read
	fmt.Println("> Read")
	if err := scenario.Read(db); err != nil {
		panic(err)
	}

	// TODO: Update

	// Delete
	fmt.Println("> Delete")
	if err := scenario.Delete(db); err != nil {
		panic(err)
	}
}
