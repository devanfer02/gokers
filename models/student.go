package models

import (
	"github.com/shopspring/decimal"
)

type Student struct {
	Model
	Name		string				`gorm:"type:varchar(155);not null" json:"name" valid:"required,type(string)"` 
	Nim 		string 				`gorm:"type:varchar(50);not null;" json:"nim" valid:"stringlength(16|16),type(string)"`
	UniModel
	Ipk 		decimal.Decimal		`gorm:"type:decimal(4,2);default:0" json:"ipk" valid:"float,type(float)"`
	Entrance	string 				`gorm:"type:varchar(100);not null" json:"entrance" valid:"required,type(string)"`
	KRS			[]KRS				`gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	KrsDetail 	KrsDetail			`gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StudentUpdate struct {
	Model
	Name		string				`gorm:"type:varchar(155);not null" json:"name" valid:"required,type(string)"` 
	Nim 		string 				`gorm:"type:varchar(50);not null;->" json:"nim" valid:"required,stringlength(16|16),type(string)"`
	UniModel
	Ipk 		decimal.Decimal		`gorm:"type:decimal(4,2);default:0" json:"ipk" valid:"required,float,type(float)"`
	Entrance	string 				`gorm:"type:varchar(100);not null" json:"entrance" valid:"required,type(string)"`
}
type StudentAuth struct {
	Nim 		string 		`json:"nim" valid:"required,numeric"`
	Password 	string 		`json:"password" valid:"required,type(string)"`
}

