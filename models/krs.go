package models

import "github.com/google/uuid"

type KRS struct {
	Model 
	StudentID 	uuid.UUID	`gorm:"foreignKey:StudentID;type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"student_id"`
	ClassID		uuid.UUID	`gorm:"foreignKey:ClassID;type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class_id"`
	Semester 	string 		`gorm:"type:varchar(100)" json:"semester" valid:"type(string)"`
	Grade 		string 		`gorm:"type:varchar(10);default:'K'" json:"grade" valid:"type(string),stringlength(1|1)"`
	Sks 		uint		`gorm:"type:integer;not null;default:0" json:"total_sks" valid:"type(integer)"`
	Student 	Student
}

type KrsDetail struct {
	Model
	StudentID	uuid.UUID 	`gorm:"foreignKey:StudentID;type:varchar(255);not null;constraint:OnUpdate:CASCADE;OnDelete:CASCADE" json:"student_id"`
	TotalSks 	uint 		`gorm:"type:integer;default:0" json:"total_sks" valid:"type(integer);range(0|24)"`	
	MaxSks 		uint 		`gorm:"type:integer;default:24" json:"max_sks" valid:"type(integer);range(0|24)"`	
}

func NewKrsDetail(id uuid.UUID, studentId uuid.UUID) *KrsDetail {
	krsDtl := &KrsDetail{}
	krsDtl.ID = id
	krsDtl.StudentID = studentId
	krsDtl.MaxSks = 24
	return krsDtl
}