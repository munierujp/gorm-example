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
)

func main() {
	db, err := gorm.Open("mysql", USER_NAME+":"+PASSWORD+"@tcp("+HOST+":"+PORT+")/"+DATABASE+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// TODO: Create

	// TODO: Read

	// TODO: Update

	// TODO: Dalete
}
