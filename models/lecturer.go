package models

type Lecturer struct {
	Model 
	Name		string	 	`gorm:"type:varchar(150);not null" json:"name" valid:"isalpha"`
	LecturerAuth
}

type LecturerAuth struct {
	Nip 		string 		`gorm:"type:varchar(150);not null" json:"nip" valid:"isalpha,required"`
	Password 	string 		`gorm:"type:varchar(150);not null" json:"password" valid:"isalphanum,required"`
}