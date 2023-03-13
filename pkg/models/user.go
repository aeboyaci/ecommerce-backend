package models

import "time"

type User struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
