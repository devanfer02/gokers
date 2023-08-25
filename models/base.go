package models

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID			uuid.UUID 	`gorm:"type:varchar(255);primaryKey" json:"id"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type UniModel struct {
	Email		string  	`gorm:"type:varchar(200);not null" json:"email" valid:"email,required,type(string)"`
	Password 	string 		`gorm:"type:varchar(255);not null" json:"password" valid:"required,stringlength(8|255),alphanum"`
	Faculty 	string		`gorm:"type:varchar(155);not null" json:"faculty" valid:"required,type(string)"`
	Major 		string		`gorm:"type:varchar(155);not null" json:"major" valid:"required,type(string)"`
}