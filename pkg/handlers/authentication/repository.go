package authentication

import (
	"ecommerce-backend/pkg/models"
	"gorm.io/gorm"
)

type repository interface {
	getUserByUsername(db *gorm.DB, username string) (models.User, error)
	save(db *gorm.DB, user models.User) error
}

type repositoryImpl struct{}

func (r repositoryImpl) getUserByUsername(db *gorm.DB, username string) (models.User, error) {
	var result models.User
	err := db.Model(&models.User{}).Where("username = ?", username).Take(&result).Error
	return result, err
}

func (r repositoryImpl) save(db *gorm.DB, user models.User) error {
	return db.Create(&user).Error
}
