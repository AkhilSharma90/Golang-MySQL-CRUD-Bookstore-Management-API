package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

const DNS = "root:root@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"

func Connect() {
	d, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
