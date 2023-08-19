package services

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	
}

func (auth *AuthService) RegisterStudent(student models.Student) res.Response {
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
	student.ID 		 = helpers.GenerateUUID()

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	fmt.Println(student.Nim)

	if err = configs.DB.Create(&student).Error; err != nil {
		return res.CreateResponseErr(status.Conflict, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "student registered to system", gin.H {
		"nim": student.Nim,
	})
}

func (auth *AuthService) LoginStudent(stdAuth models.StudentAuth) (res.Response, string) {
	if _, err := govalidator.ValidateStruct(stdAuth); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err), ""
	}

	var student models.Student

	if err := configs.DB.First(&student,"nim = ?", stdAuth.Nim).Error; err != nil {
		return res.CreateResponseErr(status.Forbidden, "invalid nim or password", nil), ""
	}

	if err := helpers.CompPassword(&student.Password, &stdAuth.Password); err != nil {
		return res.CreateResponseErr(status.Forbidden, "invalid nim or password", nil), ""
	}

	token, err := helpers.GetTokenStr(student.ID)

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err), ""
	}

	return res.CreateResponse(status.Accepted, "student successfully login", nil), token
}
