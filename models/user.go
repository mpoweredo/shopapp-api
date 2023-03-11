package models

import (
	"database/sql"
)

type User struct {
	Id          uint         `gorm:"primary_key" gorm:"autoIncrement" json:"id"`
	Username    string       `json:"username" gorm:"type:varchar(100);not null"`
	Email       string       `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password    string       `json:"password" gorm:"type:varchar(100);not null"`
	Firstname   string       `json:"firstname" gorm:"type:varchar(40)"`
	Lastname    string       `json:"lastname" gorm:"type:varchar(40)"`
	Country     string       `json:"country" gorm:"type:varchar(60)"`
	Province    string       `json:"province" gorm:"type:varchar(60)"`
	City        string       `json:"city" gorm:"type:varchar(50)"`
	Postcode    string       `json:"postcode" gorm:"type:varchar(6)"`
	Description string       `json:"description" gorm:"type:varchar(600)"`
	Photo       string       `json:"photo"`
	UpdatedAt   sql.NullTime `json:"updatedAt"`
	CreatedAt   sql.NullTime `json:"createdAt"`
}
