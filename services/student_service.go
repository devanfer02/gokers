package services

import (
	"github.com/gin-gonic/gin"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type StudentService struct {
	Db *configs.Database
}

func (stdSvc *StudentService) UpdateStudent(student *models.Student) res.Response {
	if _, err := govalidator.ValidateStruct(student); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	if stdSvc.Db.Update("id = ?", student, student.ID) == 0 {
		return res.CreateResponseErr(
			status.ServerError, 
			"failed to update data",
			fmt.Errorf("data value doesnt change or data doesnt exist"),
		)
	}

	return res.CreateResponse(status.Ok, "student data updated", gin.H {
		"record_data": student, 
	})
}