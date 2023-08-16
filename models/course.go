package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name 	string	`gorm:"varchar(155);not null" json:"name" valid:"required"`
	Code 	string 	`gorm:"varchar(8);not null" json:"code" valid:"required,stringlength(8|8)"`
}