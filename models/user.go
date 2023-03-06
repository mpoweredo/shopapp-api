package models

import (
	"database/sql"
)

type User struct {
	Id        uint         `gorm:"primary_key" gorm:"autoIncrement" json:"id"`
	Username  string       `json:"username" gorm:"type:varchar(100);not null" validate:"required,min=5,max=24"`
	Email     string       `json:"email" gorm:"type:varchar(100);unique;not null" validate:"required,email,min=6,max=48"`
	Password  string       `json:"password" gorm:"type:varchar(100);not null" validate:"required,min=5,max=24"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
	CreatedAt sql.NullTime `json:"createdAt"`
}
