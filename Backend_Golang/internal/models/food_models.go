package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name  string `json:"name" gorm:"varchar(200)"`
	Info  string `json:"info"`
	Price int    `json:"price"`
}
