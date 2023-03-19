package models

type DeliveryAddresses struct {
	Id            uint   `gorm:"primary_key" gorm:"autoIncrement" json:"id"`
	Firstname     string `gorm:"type:varchar(40)"`
	Lastname      string `gorm:"type:varchar(40)"`
	StreetAddress string `gorm:"type:varchar(60)"`
	Building      string `gorm:"type:varchar(30)"`
	City          string `gorm:"type:varchar(50)"`
	PostCode      string `gorm:"type:varchar(6)"`
	Country       string `gorm:"type:varchar(60)"`
	Province      string `gorm:"type:varchar(60)"`
	Phone         string `gorm:"type:varchar(9)"`
	UserId        uint   `gorm:"column:user_id"`
}
