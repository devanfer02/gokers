package models

type Course struct {
	Model 
	CourseName 	string	 	`gorm:"type:varchar(155);not null" json:"name" valid:"required,alpha"`
	CourseCode 	string 	 	`gorm:"type:varchar(8);not null" json:"code" valid:"required,alpha,stringlength(8|8)"`
	Sks			int		 	`gorm:"type:integer;default:1" json:"sks" valid:"required,int"`
	Major 		string 		`gorm:"type:varchar(100);not null" json:"type" valid:"required,string"`
	Class   	[]Class  	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}