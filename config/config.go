package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// initialize sqlite database
	db, err = InitializeSQlite()
	if err != nil {
		return fmt.Errorf("Error initializing sqlite database: %v", err)
	}

	return nil
}

func GetSQlite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize logger
	logger = NewLogger(p)
	return logger
}
