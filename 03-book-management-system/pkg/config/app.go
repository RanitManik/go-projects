package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(
		"mysql",
		"bookuser:bookpass@tcp(127.0.0.1:3306)/bookdb?charset=utf8mb4&parseTime=True&loc=Local",
	)
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("mySQL database successfully connected...")
}

func GetDB() *gorm.DB {
	return db
}
