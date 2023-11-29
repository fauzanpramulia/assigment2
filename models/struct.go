package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Item_code string `json:"itemCode" form:"itemCode" gorm:"not null;varchar(200)"`
	Description string `json:"description" form:"description" gorm:"not null;varchar(200)"`
	Quantity int `json:"quantity" form:"quantity" gorm:"not null"`
	Order_id int `gorm:"not null"`
}

type Order struct {
	gorm.Model
	Customer_name string `json:"customer_name" form:"customer_name" query:"customer_name" gorm:"not null;varchar(200)"`
	Items         []Item
	Ordered_at    time.Time `gorm:"not null"`
}

type Response struct {
	Messages string
	Success  bool
	Data     interface{}
}