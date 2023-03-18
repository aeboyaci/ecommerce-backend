package address

import (
	"ecommerce-backend/pkg/models"
	"gorm.io/gorm"
)

type repository interface {
	save(db *gorm.DB, address models.Address) error
}

type repositoryImpl struct{}

func (r repositoryImpl) save(db *gorm.DB, address models.Address) error {
	return db.Create(&address).Error
}
