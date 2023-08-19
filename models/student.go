package models

import (
	"github.com/shopspring/decimal"
)

type Student struct {
	Model
	Name		string				`gorm:"type:varchar(155);not null" json:"name" valid:"required,type(string)"` 
	Nim 		string 				`gorm:"type:varchar(50);not null;" json:"nim" valid:"stringlength(16|16),type(string)"`
	Email		string  			`gorm:"type:varchar(200);not null" json:"email" valid:"email,required,type(string)"`
	Password 	string 				`gorm:"type:varchar(255);not null" json:"password" valid:"required,stringlength(8|255),alphanum"`
	Major 		string				`gorm:"type:varchar(155);not null" json:"major" valid:"required,type(string)"`
	Faculty 	string				`gorm:"type:varchar(155);not null" json:"faculty" valid:"required,type(string)"`
	Ipk 		decimal.Decimal		`gorm:"type:decimal(4,2);default:0" json:"ipk" valid:"float,type(float)"`
	Entrance	string 				`gorm:"type:varchar(100);not null" json:"entrance" valid:"required,type(string)"`
	KRS			[]KRS				`gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StudentUpdate struct {
	Model
	Name		string				`gorm:"type:varchar(155);not null" json:"name" valid:"required,type(string)"` 
	Nim 		string 				`gorm:"type:varchar(50);not null;->" json:"nim" valid:"required,stringlength(16|16),type(string)"`
	Email		string  			`gorm:"type:varchar(200);not null" json:"email" valid:"email,required,type(string)"`
	Password 	string 				`gorm:"type:varchar(255);not null" json:"password" valid:"required,stringlength(8|255),alphanum"`
	Major 		string				`gorm:"type:varchar(155);not null" json:"major" valid:"required,type(string)"`
	Faculty 	string				`gorm:"type:varchar(155);not null" json:"faculty" valid:"required,type(string)"`
	Ipk 		decimal.Decimal		`gorm:"type:decimal(4,2);default:0" json:"ipk" valid:"required,float,type(float)"`
	Entrance	string 				`gorm:"type:varchar(100);not null" json:"entrance" valid:"required,type(string)"`
}

type StudentAuth struct {
	Nim 		string 		`json:"nim" valid:"required,numeric"`
	Password 	string 		`json:"password" valid:"required,type(string)"`
}
