package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var gormDB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to the database!")
	}

	db.AutoMigrate(&Book{})

	gormDB = db
}
