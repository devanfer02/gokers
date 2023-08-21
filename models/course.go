package models

type Course struct {
	Model 
	CourseName 	string	 	`gorm:"type:varchar(155);not null" json:"name" valid:"required,type(string)"`
	CourseCode 	string 	 	`gorm:"type:varchar(8);not null" json:"code" valid:"required,alphanum,stringlength(8|8)"`
	Sks			int		 	`gorm:"type:integer;default:1" json:"sks" valid:"required,int"`
	Type		string 		`gorm:"type:varchar(100);not null" json:"type" valid:"required,type(string)"`
	Faculty 	string 		`gorm:"type:varchar(100);default:''" json:"faculty" valid:"type(string)"`
	Major 		string 		`gorm:"type:varchar(100);default:''" json:"major" valid:"type(string)"`
	Class   	[]Class  	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class,omitempty"`
}