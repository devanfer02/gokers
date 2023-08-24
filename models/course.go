package models

type Course struct {
	Model 
	CourseName 			string	 			`gorm:"type:varchar(155);not null" json:"name" valid:"required,type(string)"`
	CourseCode 			string 	 			`gorm:"type:varchar(8);not null;unique" json:"code" valid:"required,alphanum,stringlength(8|8)"`
	Sks					int		 			`gorm:"type:integer;default:1" json:"sks" valid:"required,int"`
	Type				string 				`gorm:"type:varchar(100);not null" json:"type" valid:"required,type(string)"`
	Faculty 			string 				`gorm:"type:varchar(100);default:''" json:"faculty" valid:"type(string)"`
	Major 				string 				`gorm:"type:varchar(100);default:''" json:"major" valid:"type(string)"`
	Class   			[]Class  			`gorm:"foreignKey:CourseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"class,omitempty"`
	UpperCourse			CoursePrequisites 	`gorm:"foreignKey:UpperID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"upper,omitempty"`
	CoursePrequisites 	[]CoursePrequisites	`gorm:"foreignKey:PrequisiteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"prequisites,omitempty"`
}

type CoursePrequisites struct {
	Model
	UpperID 		string	 `gorm:"foreignKey:UpperID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"course_id"`
	PrequisiteID	string	 `gorm:"foreignKey:PrequisiteID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE" json:"prequisite_id"`
}