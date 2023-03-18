package address

import (
	"ecommerce-backend/pkg/common/database"
	"ecommerce-backend/pkg/models"
)

type controller interface {
	addNewAddress(userId string, address models.Address) error
}

type controllerImpl struct {
	repository repository
}

func newController() controller {
	return controllerImpl{
		repository: repositoryImpl{},
	}
}

func (c controllerImpl) addNewAddress(userId string, address models.Address) error {
	address.UserID = userId
	return c.repository.save(database.GetInstance(), address)
}
