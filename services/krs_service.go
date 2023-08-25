package services

import (
	//"github.com/gin-gonic/gin"
	"fmt"

	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers"

	//"github.com/devanfer02/gokers/helpers"
	//"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type KrsService struct {
	Db *configs.Database
}

func NewKrsService(db *configs.Database) *KrsService {
	return &KrsService{Db: db}
}

func (krsSvc *KrsService) GetKrs(krs *[]models.KRS, params ...interface{}) res.Response {
	if err := krsSvc.Db.FindAllCondition("student_id = ?", &krs, params); err != nil {
		return res.CreateResponseErr(
			status.ServerError, 
			"internal server error", 
			err, 
		)
	}

	return res.CreateResponse(status.Ok, "succesfully fetch krs", krs)
}

func (krsSvc *KrsService) AddClass(krs *models.KRS, params ...interface{}) res.Response {
	student := &models.Student{}
	class := &models.Class{}
	detail := &models.KrsDetail{}

	if err := krsSvc.Db.FirstByPK(student, krs.StudentID); err != nil {
		return res.CreateResponseErr(
			status.NotFound, 
			"student not found", 
			err, 
		)
	}

	if err := krsSvc.Db.PreloadByPK([]string{"Course"}, class, krs.ClassID); err != nil {
		return res.CreateResponseErr(
			status.NotFound, 
			"class not found", 
			err, 
		)
	}

	if err := krsSvc.Db.Find("student_id = ?", detail, krs.StudentID); err != nil {
		return res.CreateResponseErr(
			status.NotFound, 
			"class not found", 
			err, 
		)
	}

	if class.Quota >= 40 {
		return res.CreateResponseErr(
			status.TooMany, 
			"class quota has reached maximum",
			fmt.Errorf("class quota is full"),
		)
	}

	if detail.TotalSks + class.Course.Sks > detail.MaxSks {
		return res.CreateResponseErr(
			status.TooMany, 
			"maximum sks reached",
			fmt.Errorf("cannot add class, overflow max sks"),
		)
	}

	krs.Semester = helpers.GetCurrentSemester()
	detail.TotalSks = detail.TotalSks + class.Course.Sks
	class.Quota = class.Quota + 1

	if err := krsSvc.Db.Create(krs); err != nil {
		return res.CreateResponseErr(
			status.ServerError, 
			"internal server error", 
			err, 
		)
	}

	if row := krsSvc.Db.Update("student_id = ?", detail, krs.StudentID); row == 0 {
		return res.CreateResponseErr(
			status.ServerError, 
			"failed to update data", 
			fmt.Errorf("failed to add total sks to krs_detail table"), 
		)
	}

	if row := krsSvc.Db.UpdateByPK(class, class.ID); row == 0 {
		return res.CreateResponseErr(
			status.ServerError, 
			"failed to update data", 
			fmt.Errorf("failed to add total sks to krs_detail table"), 
		)
	}

	return res.CreateResponse(status.Ok, "student's krs updated", nil)
}

func (krsSvc *KrsService) RemoveClass(krs *models.KRS) res.Response {
	if rows := krsSvc.Db.Delete("student_id = ? AND class_id = ?", krs, krs.StudentID, krs.ClassID); rows == 0 {
		return res.CreateResponseErr(
			status.ServerError, 
			"failed to delete data", 
			fmt.Errorf("failed to delete krs"), 
		)
	}

	return res.CreateResponse(status.Ok, "successfully delete krs", nil)
}