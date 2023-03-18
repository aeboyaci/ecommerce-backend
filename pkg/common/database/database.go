package database

import (
	"ecommerce-backend/pkg/common/env"
	log "ecommerce-backend/pkg/common/logger"
	"ecommerce-backend/pkg/models"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Initialize() error {
	var err error
	db, err = gorm.Open(postgres.Open(env.DB_URL), &gorm.Config{
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
		log.GetInstance().Error(err.Error())
		return errors.New("cannot migrate User table")
	}

	err = db.AutoMigrate(&models.Address{})
	if err != nil {
		log.GetInstance().Error(err.Error())
		return errors.New("cannot migrate Address table")
	}

	log.GetInstance().Info("migrations run successfully")
	return nil
}
