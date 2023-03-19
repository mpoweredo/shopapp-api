package models

import (
	"database/sql"
)

type User struct {
	Id                uint   `gorm:"primary_key" gorm:"autoIncrement" json:"id"`
	Username          string `gorm:"type:varchar(100);not null"`
	Email             string `gorm:"type:varchar(100);unique;not null"`
	Password          string `gorm:"type:varchar(100);not null"`
	Firstname         string `gorm:"type:varchar(40)"`
	Lastname          string `gorm:"type:varchar(40)"`
	Country           string `gorm:"type:varchar(60)"`
	Province          string `gorm:"type:varchar(60)"`
	City              string `gorm:"type:varchar(50)"`
	Postcode          string `gorm:"type:varchar(6)"`
	Description       string `gorm:"type:varchar(600)"`
	Photo             string
	DeliveryAddresses []DeliveryAddresses
	UpdatedAt         sql.NullTime
	CreatedAt         sql.NullTime
}
