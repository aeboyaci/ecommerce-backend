package address

import (
	"ecommerce-backend/pkg/models"
	"gorm.io/gorm"
)

type repository interface {
	save(db *gorm.DB, address models.Address) error
	getAllAddresses(db *gorm.DB, userId string) ([]AddressResponseDTO, error)
	updateAddressById(db *gorm.DB, userId string, addressId string, address models.Address) error
}

type repositoryImpl struct{}

func (r repositoryImpl) save(db *gorm.DB, address models.Address) error {
	return db.Create(&address).Error
}

func (r repositoryImpl) getAllAddresses(db *gorm.DB, userId string) ([]AddressResponseDTO, error) {
	var result []AddressResponseDTO
	err := db.Model(&models.Address{}).Where("user_id = ?", userId).Find(&result).Error
	return result, err
}

func (r repositoryImpl) updateAddressById(db *gorm.DB, userId string, addressId string, address models.Address) error {
	return db.
		Model(&address).
		Where("user_id = ? AND id = ?", userId, addressId).
		Updates(&address).
		Error
}
