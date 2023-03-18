package address

import (
	"ecommerce-backend/pkg/models"
	"gorm.io/gorm"
)

type repository interface {
	save(db *gorm.DB, address models.Address) error
	getAllAddresses(db *gorm.DB, userId string) ([]models.Address, error)
}

type repositoryImpl struct{}

func (r repositoryImpl) save(db *gorm.DB, address models.Address) error {
	return db.Create(&address).Error
}

func (r repositoryImpl) getAllAddresses(db *gorm.DB, userId string) ([]models.Address, error) {
	var result []models.Address
	err := db.Model(&models.Address{}).Where("user_id = ?", userId).Find(&result).Error
	return result, err
}
