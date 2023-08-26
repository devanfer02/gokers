package models

type Lecturer struct {
	Model 
	Name		string	 	`gorm:"type:varchar(150);not null" json:"name" valid:"type(string),required"`
	Ndn 		string 		`gorm:"type:varchar(150);not null" json:"ndn" valid:"type(string)"`
	UniModel
	Class	  	*[]Class  	`gorm:"foreignKey:LecturerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"classes,omitempty"`
}

type LecturerOptional struct {
	Model 
	Name		string	 	`gorm:"type:varchar(150);not null" json:"name" valid:"type(string)"`
	Ndn 		string 		`gorm:"type:varchar(150);not null" json:"ndn" valid:"type(string)"`
	UniModel
	Class	  	[]Class  	`gorm:"foreignKey:LecturerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}