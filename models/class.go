package models

import "github.com/google/uuid"

type Class struct {
	Model
	CourseID	uuid.UUID	`gorm:"foreignKey:CourseID;type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"course_id"`
	LecturerID 	uuid.UUID 	`gorm:"foreignKey:LecturerID;type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"lecturer_id"`
	ClassName  	string		`gorm:"type:varchar(100);not null" json:"class_name" valid:"required"`
	ClassRoom  	string		`gorm:"type:varchar(100);not null" json:"class_room" valid:"required"`
	KRS        	[]KRS		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}