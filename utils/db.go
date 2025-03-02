package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tests.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
