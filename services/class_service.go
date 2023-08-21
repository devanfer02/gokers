package services

import (
	//"fmt"

	//"github.com/gin-gonic/gin"
	//"github.com/asaskevich/govalidator"
	"github.com/devanfer02/gokers/configs"
	//"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	//"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type ClassService struct {
	Db *configs.Database
}

func (classSvc *ClassService) RegisterClass(class models.Class) res.Response {
	var course models.Course

	if err := classSvc.Db.Find("id = ?", course, course.ID); err != nil {
		return res.CreateResponse(status.Accepted, "ok", nil)
	}

	return res.CreateResponse(status.Accepted, "ok", nil)
}