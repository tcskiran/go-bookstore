package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

// need to return db so that other pacakges can interact with the database

var (
	db *gorm.DB
)

func Connect() {
	// dsn := "root:password@tcp(172.22.0.2:3306)/dvba?charset=utf8mb4&parseTime=True&loc=Local"
	// d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// d, err := gorm.Open("mysql", "root:password@(172.22.0.2:3306)/dvba?/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "root:password@(172.22.0.2:3306)/dvba?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
