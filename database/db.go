package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init - initialize database
func Init() {

	db, err := gorm.Open(sqlite.Open("voodoo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate()

}

// DbManager - get database instance
func DbManager() *gorm.DB {
	return db
}
