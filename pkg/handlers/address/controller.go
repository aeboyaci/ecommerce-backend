package address

import (
	"ecommerce-backend/pkg/common/database"
	"ecommerce-backend/pkg/models"
	"errors"
	"strings"
)

type controller interface {
	addNewAddress(userId string, address models.Address) error
	getAllAddresses(userId string) ([]AddressResponseDTO, error)
	updateAddressById(userId string, addressId string, address models.Address) error
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
	if !strings.HasPrefix(address.PhoneNumber, "+90") {
		return errors.New("invalid phone number given")
	}

	address.UserID = userId
	return c.repository.save(database.GetInstance(), address)
}

func (c controllerImpl) getAllAddresses(userId string) ([]AddressResponseDTO, error) {
	return c.repository.getAllAddresses(database.GetInstance(), userId)
}

func (c controllerImpl) updateAddressById(userId string, addressId string, address models.Address) error {
	if address.PhoneNumber != "" && !strings.HasPrefix(address.PhoneNumber, "+90") {
		return errors.New("invalid phone number given")
	}
	return c.repository.updateAddressById(database.GetInstance(), userId, addressId, address)
}
