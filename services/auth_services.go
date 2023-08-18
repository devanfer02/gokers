package services

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type AuthService struct {
	
}

func (auth *AuthService) RegisterStudent (student models.Student) res.Response {
	if _, err := govalidator.ValidateStruct(student); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	var studentdb models.Student
	count := int64(0)

	configs.DB.
		Model(&studentdb).
		First(&studentdb, "email = ?", student.Email). 
		Count(&count)

	if count > 0 {
		return res.CreateResponseErr(status.Conflict, "email already in use", fmt.Errorf("email is already used"))
	}

	hashed, err := helpers.HashPassword(student.Password)

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	student.Password = hashed
	student.Nim, err = helpers.CreateNIM(student.Major, student.Faculty, student.Entrance)

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	err = configs.DB.Create(&student).Error 

	if err != nil {
		return res.CreateResponseErr(status.Conflict, "internal server error", err)
	}

	return res.CreateResponseOk(status.Ok, "student registered to system", nil)
}