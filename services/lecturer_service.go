package services

import (
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type LecturerService struct {
	Db *configs.Database
}

func NewLecturerService(db *configs.Database) *LecturerService {
	return  &LecturerService{Db: db}
}

func (lctSvc *LecturerService) GetLecturers(lecturers *[]models.Lecturer) res.Response {
	if err := lctSvc.Db.FindAll(lecturers); err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "successfully fetch lecturers", lecturers) 
}