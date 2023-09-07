package services

import (
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/models/constructors"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type AdminService struct {
	Db *configs.Database
}

func NewAdminService(db *configs.Database) *AdminService {
	return &AdminService{Db: db}
}

func (admSvc *AdminService) AddStudent(student *models.Student, class *models.Class) res.Response {
	if err := admSvc.Db.FindFirst("id = ?", student, student.ID); err != nil {
		return res.CreateResponseErr(status.NotFound, "student not found", err)
	}

	if err := admSvc.Db.PreloadByPK([]string{"Course"}, class, class.ID); err != nil {
		return res.CreateResponseErr(status.NotFound, "student not found", err)
	}

	krs := constructor.NewKRS(student.ID, class.ID, class.Course.Sks)

	if err := admSvc.Db.Create(krs); err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "successfully insert student to class", krs);
}

func (admSvc *AdminService) RemoveStudent(student *models.Student, class *models.Class) res.Response {
	krs := &models.KRS{}

	if row := admSvc.Db.Delete("student_id = ? AND class_id = ?", krs, student.ID, class.ID); row == 0 {
		return res.CreateResponseErr(status.NotFound, "krs not found", nil)
	}

	return res.CreateResponse(status.Ok, "successfully remove student from class", nil);
}