package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type CourseService struct {
	Db *configs.Database
}

func (courseSvc *CourseService) RegisterCourse(course models.Course) res.Response {
	if _, err := govalidator.ValidateStruct(course); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	count := courseSvc.Db.Count("course_code", models.Course{}, course.CourseCode)	

	if count > 0 {
		return res.CreateResponseErr(status.Conflict, "course code conflicted", fmt.Errorf("course code already exist"))
	}

	if err := helpers.CheckTypeExist(&course.Type, &course.Faculty, &course.Major); err != nil {
		return res.CreateResponseErr(status.BadRequest, "types or faculty or major doesnt exist", err)
	}

	course.ID = helpers.GenerateUUID()
	

	if err := courseSvc.Db.Create(&course); err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "course registered to system", gin.H{
		"record_data": course,
	})
}

func (courseSvc *CourseService) GetCourses(course []models.Course, queries []string) res.Response {
	var err error 

	if queries[0] == "" && queries[1] == "" && queries[2] == ""{ 
		err = courseSvc.Db.FindAll(&course);
	} else {
		var query, param string
		if queries[0] != ""  {
			query = "type = ?"
			param = queries[0]
		}
		if queries[1] != "" {
			query = "faculty = ?"
			param = queries[1]
		}

		if  queries[2] != "" {
			query = "major = ?"
			param = queries[2]
		}

		err = courseSvc.Db.Find(query, &course, param)
	}

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}	

	return res.CreateResponse(status.Ok, "sucessfully fetch course", course)
}

func (courseSvc *CourseService) GetCourse(course models.Course) res.Response {
	if err := courseSvc.Db.Find("id = ?", &course, course.ID); err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "sucessfully fetch course", course)
}

func (courseSvc *CourseService) UpdateCourse(course models.Course) res.Response {
	if _, err := govalidator.ValidateStruct(course); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	if err := helpers.CheckTypeExist(&course.Type, &course.Faculty, &course.Major); err != nil {
		return res.CreateResponseErr(status.BadRequest, "types or faculty or major doesnt exist", err)
	}	

	if courseSvc.Db.Update("id = ?", &course, course.ID) == 0 {
		return res.CreateResponseErr(status.ServerError, "failed to update data", fmt.Errorf("data value doesnt change probably"))
	}

	return res.CreateResponse(status.Ok, "course data updated", gin.H {
		"record_data": course,
	})
}

func (courseSvc *CourseService) DeleteService(course models.Course) res.Response {
	if courseSvc.Db.Delete("id = ?", course, course.ID) == 0 {
		return res.CreateResponseErr(status.ServerError, "failed to delete data", fmt.Errorf("data didnt exist"))
	}

	return res.CreateResponse(status.Ok, "successfully delete course", gin.H {
		"deleted_record_data": course,
	})
}