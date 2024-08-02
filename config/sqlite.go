package config

import (
	"montravail/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./database/events.db"

	// Check if existing database
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating ...")
		// Create database file and directory
		err = os.MkdirAll("./database", os.ModePerm)
		if err != nil {
			logger.Errorf("Error: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create database and connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite open error: %v", err)
		return nil, err
	}
	// Migrate and schema
	err = db.AutoMigrate(&schemas.Event{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	// Return database
	return db, nil

}
