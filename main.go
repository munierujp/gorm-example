package main

import (
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
		panic("failed to connect database")
	}
	defer db.Close()

	// TODO: Create

	// TODO: Read

	// TODO: Update

	// TODO: Dalete
}
