package models

import "github.com/google/uuid"

type Class struct {
	Model
	ClassModel
	Quota     	uint    `gorm:"type:integer;default:40;not null" json:"quota" valid:"type(uint),range(5|100)"`
	Current		uint	`gorm:"type:integer;default:0;not null" json:"current" valid:"type(uint),range(0|100)"`	
	KRS       	*[]KRS   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"krs,omitempty"`
}

type ClassRegister struct {
	CourseID  	uuid.UUID 	`gorm:"type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"course_id" valid:"required"`
	LecturerID  uuid.UUID 	`gorm:"type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"lecturer_id" valid:"required"`
	ClassName 	string    	`gorm:"type:varchar(100);not null" json:"class_name" valid:"required"`
	ClassRoom 	string    	`gorm:"type:varchar(100);not null" json:"class_room" valid:"required"`
	Quota     	uint    `gorm:"type:integer;default:40;not null" json:"quota" valid:"type(uint),range(5|100)"`
	Current		uint	`gorm:"type:integer;default:0;not null" json:"current" valid:"type(uint),range(0|100)"`	
}

type ClassModel struct {
	Course 		*Course 	`json:"course_info,omitempty"` 
	Lecturer 	*Lecturer	`json:"lecturer_info,omitempty"`
	CourseID  	uuid.UUID 	`gorm:"type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"course_id" valid:"required"`
	LecturerID  uuid.UUID 	`gorm:"type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"lecturer_id" valid:"required"`
	ClassName 	string    	`gorm:"type:varchar(100);not null" json:"class_name" valid:"required"`
	ClassRoom 	string    	`gorm:"type:varchar(100);not null" json:"class_room" valid:"required"`
}


