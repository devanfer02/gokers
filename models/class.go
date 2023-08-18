package models

import "github.com/google/uuid"

type Class struct {
	Course
	LecturerID 	uuid.UUID 	`gorm:"type:uuid;not null" json:"lecturer_id"`
	ClassName  	string		`gorm:"type:varchar(100);not null" json:"class_name" valid:"required"`
	ClassRoom  	string		`gorm:"type:varchar(100);not null" json:"class_room" valid:"required"`
	KRS        	[]KRS		`gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}