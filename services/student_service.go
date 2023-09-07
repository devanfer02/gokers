package services

import (
	//"github.com/gin-gonic/gin"
	//"fmt"

	"github.com/devanfer02/gokers/configs"
	//"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type StudentService struct {
	Db *configs.Database
}

func (stdSvc *StudentService) UpdateStudent(student *models.Student) res.Response {
	return res.CreateResponse(status.Ok, "OK", nil);
}