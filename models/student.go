package models

import (
	"github.com/shopspring/decimal"
)

type Student struct {
	Model
	StudentAuth 
	Name		string				`gorm:"type:varchar(155);not null" json:"name" valid:"required,alpha"` 
	Major 		string				`gorm:"type:varchar(155);not null" json:"major" valid:"required,alpha"`
	Faculty 	string				`gorm:"type:varchar(155);not null" json:"faculty" valid:"required,alpha"`
	Email		string  			`gorm:"type:varchar(200);not null" json:"email" valid:"email,required"`
	Ipk 		decimal.Decimal		`gorm:"type:decimal(4,2);default:0" json:"ipk" valid:"float"`
	KRS			[]KRS				`gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StudentAuth struct {
	Nim 		string 		`gorm:"type:varchar(16);not null" json:"nim" valid:"required,stringlength(16|16)"`
	Password 	string 		`gorm:"type:varchar(255);not null" json:"password" valid:"required,stringlength(8|255),isalphanum"`
}