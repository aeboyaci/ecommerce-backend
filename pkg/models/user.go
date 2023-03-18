package models

type User struct {
	Base
	Username  string `json:"username" gorm:"uniqueIndex"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
