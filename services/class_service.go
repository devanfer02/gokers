package services

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type ClassService struct {
	Db *configs.Database
}

func NewClassService(db *configs.Database) *ClassService {
	return &ClassService{Db : db}
}

func (classSvc *ClassService) GetParticipants(class *models.Class, student *models.Student) res.Response {
	var krs *[]models.KRS
	var students []models.Student

	if err := classSvc.Db.FirstByPK(class, class.ID); err != nil {
		return res.CreateResponseErr(status.NotFound, "class not found", err)
	}

	if err := classSvc.Db.PreloadByCondition([]string{"Student"}, krs, "class_id = ?" ,class.ID); err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	forbidden := true

	for _, unit := range *krs {
		if unit.Student.ID == student.ID {
			forbidden = false
		}
		students = append(students, unit.Student)
	}

	if forbidden {
		return res.CreateResponse(status.Forbidden, "forbidden", nil)
	}

	return res.CreateResponse(status.Ok, "successfully fetch participants", gin.H {
		"class_info": class, 
		"participants": students,
	})
}

func (classSvc *ClassService) GetClasses(classes *[]models.Class, queries []string, student *models.Student, ok bool) res.Response {
	var err error 

	if queries[0] != "" {
		err = classSvc.Db.PreloadByCondition([]string{"Course", "Lecturer"}, classes, "course_code = ?", queries[0]);
	} else {
		err = classSvc.Db.PreloadMany([]string{"Course", "Lecturer"}, classes);
	}

	if err != nil {
		return res.CreateResponseErr(status.ServerError,
			"internal server error", 
			err,
		)
	}

	if ok {
		var filtered []models.Class 
		for _, class := range *classes {
			if class.Course.Major == student.Major {
				filtered = append(filtered, class)
			}
		}

		*classes = filtered
	}

	return res.CreateResponse(status.Ok, "successfully fetch classes", classes)
}

func (classSvc *ClassService) GetClass(class *models.Class, student *models.Student, ok bool) res.Response {
	if err := classSvc.Db.PreloadByPK([]string{"Course","Lecturer"}, class, class.ID); err != nil {
		return res.CreateResponseErr(status.BadRequest, "foreign entity assosiated not found body", err)
	}

	if ok && class.Course.Major != student.Major {
		return res.CreateResponseErr(status.NotFound, "class not found", nil)
	}

	return res.CreateResponse(status.Ok, "succesfully fetch class", class)
}

func (classSvc *ClassService) RegisterClass(class *models.Class) res.Response {
	var course models.Course
	var lecturer models.Lecturer

	if _, err := govalidator.ValidateStruct(class); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	if count := classSvc.Db.Count("class_name", models.Class{}, class.ClassName); count > 0 {
		return res.CreateResponseErr(status.Conflict, "class with that name already exist", nil)
	}

	if err := classSvc.Db.FirstByPK(&course, class.CourseID); err != nil {
		return res.CreateResponseErr(status.NotFound, "course_id not found", err)
	}

	if err := classSvc.Db.FirstByPK(&lecturer, class.LecturerID); err != nil { 
		return res.CreateResponseErr(status.NotFound, "lecturer_id not found", err)
	}

	class.ID = helpers.GenerateUUID()

	if err := classSvc.Db.Create(class); err != nil {
		return res.CreateResponseErr(status.ServerError, "internal sever error", err)
	}

	return res.CreateResponse(status.Ok, "class registered to system", gin.H {
		"record_data": class,
	})
}

func (classSvc *ClassService) UpdateClass(class *models.Class) res.Response {
	if _, err := govalidator.ValidateStruct(class); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	if classSvc.Db.Update("id = ?", class, class.ID) == 0 {
		return res.CreateResponseErr(
			status.ServerError, 
			"failed to update data",
			fmt.Errorf("data value doesnt change or data doesnt exist"),
		)
	}

	return res.CreateResponse(status.Ok, "class data updated", gin.H {
		"record_data": class, 
	})
}

func (classSvc *ClassService) DeleteClass(class *models.Class) res.Response {
	if classSvc.Db.Delete("id = ?", class, class.ID) == 0 {
		return res.CreateResponseErr(
			status.ServerError,
			"failed to delete data",
			fmt.Errorf("data didnt exist"),
		)
	}

	return res.CreateResponse(status.Ok, "successfully delete class", nil)
}