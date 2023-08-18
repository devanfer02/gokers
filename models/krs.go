package models

import "github.com/google/uuid"

type KRS struct {
	Model 
	StudentID 	uuid.UUID	`gorm:"type:uuid" json:"student_id"`
	ClassID		uuid.UUID	`gorm:"type:uuid" json:"class_id"`
}