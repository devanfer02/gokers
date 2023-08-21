package models

type Admin struct {
	Model 
	Username	string 		`gorm:"type:varchar(200);not null" json:"username" valid:"required,alphanum,stringlength(8|200)"`
	Email		string  	`gorm:"type:varchar(200);not null" json:"email" valid:"email,required,type(string)"`
	Password 	string 		`gorm:"type:varchar(255);not null" json:"password" valid:"required,stringlength(8|255),alphanum"`
}