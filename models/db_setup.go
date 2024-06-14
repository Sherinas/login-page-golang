package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDatabase(db *gorm.DB) {
	DB = db
}

func DatabaseSetup() *gorm.DB {
	dsn := "host=localhost port=5432 user=sherinascdlm password=admin123 dbname=giorm sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}
	return db
}
