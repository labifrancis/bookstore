package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "username:password@/tablename?charset=utf8&parseTime=True&loc=Local?")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}

// The purpose of this file is to help us return a db variable that will help us talk to the database
