package database

import (
	"ecommerce-backend/pkg/models"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var db *gorm.DB

func Initialize() error {
	connectionString, hasEnvironmentVariable := os.LookupEnv("DB_URL")
	if !hasEnvironmentVariable {
		return errors.New("DB_URL environment variable has not set")
	}

	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("cannot connect to the database. Error: %s", err.Error())
	}

	return autoMigrate()
}

func GetInstance() *gorm.DB {
	return db
}

func autoMigrate() error {
	var err error

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return errors.New("cannot migrate User table")
	}

	return nil
}
