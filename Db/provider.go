package Db

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("checkr.db"), &gorm.Config{})
	Helpers.Log(err, "Initialising Db")

	// migrate Schemas
	migrateSchemas(db)

	return db
}

func migrateSchemas(db *gorm.DB) {
	err := db.AutoMigrate(&Models.Event{}, &Models.User{}, &Models.Category{}, &Models.EventList{}, &Models.Priority{}, &Models.Notification{})
	Helpers.Log(err, "Migrating Events")
}
