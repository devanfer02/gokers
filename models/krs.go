package models

import "github.com/google/uuid"

type KRS struct {
	Model 
	StudentID 	uuid.UUID	`gorm:"foreignKey:StudentID;type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"student_id"`
	ClassID		uuid.UUID	`gorm:"foreignKey:ClassID;type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class_id"`
	Sks 		uint		`gorm:"type:integer;not null;default:0" json:"total_sks" valid:"type(integer)"`
}