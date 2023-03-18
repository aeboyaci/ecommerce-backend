package models

import (
	"errors"
	"gorm.io/gorm"
	"strings"
)

type Address struct {
	Base
	UserID      string `json:"-" gorm:"index"`
	User        User   `json:"-" gorm:"foreignKey:UserID"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	PhoneNumber string `json:"phoneNumber"`
}

func (address *Address) BeforeSave(tx *gorm.DB) error {
	if !strings.HasPrefix(address.PhoneNumber, "+90") {
		return errors.New("invalid phone number given")
	}

	return nil
}
