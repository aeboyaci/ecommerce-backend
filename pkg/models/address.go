package models

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
