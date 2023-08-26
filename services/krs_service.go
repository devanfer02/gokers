package services

import (
	//"github.com/gin-gonic/gin"
	"fmt"

	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers"
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

	if class.Current >= class.Quota {
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

	if count := krsSvc.Db.CountJoins(
			"JOIN classes ON krs.class_id = classes.ID",
			[]string{"classes.course_id = ?", "student_id = ?"}, 
			krs,
			[]interface{}{class.CourseID, krs.StudentID},
		); count > 0 {
		return res.CreateResponseErr(
			status.Conflict, 
			"course already exist",
			fmt.Errorf("cannot add class, course already exist"),
		)
	}

	krs.ID = helpers.GenerateUUID()
	krs.Sks = class.Course.Sks
	krs.Semester = helpers.GetCurrentSemester()
	detail.TotalSks = detail.TotalSks + class.Course.Sks
	class.Current = class.Current + 1

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
	detail := &models.KrsDetail{}
	class := &models.Class{}

	if err := krsSvc.Db.Find("student_id = ?", detail, krs.StudentID); err != nil {
		return res.CreateResponseErr(
			status.NotFound, 
			"class not found", 
			err, 
		)
	}

	if err := krsSvc.Db.Find("student_id = ? AND class_id = ?", krs, krs.StudentID, krs.ClassID); err != nil {
		return res.CreateResponseErr(
			status.NotFound, 
			"failed to find related data id", 
			err,
		)	
	}

	detail.TotalSks = detail.TotalSks - krs.Sks 
	hack := int64(class.Current) - 1
	if hack < 0 { hack = 0 }
	class.Current = uint(hack)
	class.ID = krs.ClassID
	

	if row := krsSvc.Db.UpdateMapByCondition(
		"student_id = ?", 
		detail,
		map[string]interface{}{"TotalSks": detail.TotalSks},
		krs.StudentID,
		); row == 0 {
		return res.CreateResponseErr(
			status.ServerError, 
			"failed to update data", 
			fmt.Errorf("failed to reduce total sks in krs_detail table"), 
		)
	}

	if row := krsSvc.Db.UpdateMapByCondition(
		"id = ?",
		class, 
		map[string]interface{}{"Current": class.Current},	
		krs.ClassID,
		); row == 0 {
		return res.CreateResponseErr(
			status.ServerError, 
			"failed to update data", 
			fmt.Errorf("failed to reduce class participants in class table"), 
		)
	}

	if rows := krsSvc.Db.Delete("student_id = ? AND class_id = ?", krs, krs.StudentID, krs.ClassID); rows == 0 {
		return res.CreateResponseErr(
			status.ServerError, 
			"failed to delete data", 
			fmt.Errorf("failed to delete krs"), 
		)
	}

	return res.CreateResponse(status.Ok, "successfully delete krs", nil)
}