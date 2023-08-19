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

}

func (self *CourseService) RegisterCourse(course models.Course) res.Response {
	if _, err := govalidator.ValidateStruct(course); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	count := helpers.DbCount("course_code", models.Course{}, course.CourseCode)	

	if count > 0 {
		return res.CreateResponseErr(status.Conflict, "course code conflicted", fmt.Errorf("course code already exist"))
	}

	course.ID = helpers.GenerateUUID()

	if err := configs.DB.Create(&course).Error; err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "course registered to system", gin.H{
		"record_data": course,
	})
}

func (self *CourseService) GetCourses(course []models.Course, typequery string) res.Response {
	var err error 

	if typequery == "" {
		err = configs.DB.Find(&course).Error;
	} else {
		err = configs.DB.Where("type = ?", typequery).Find(&course).Error;
	}

	if err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}	

	return res.CreateResponse(status.Ok, "sucessfully fetch course", course)
}

func (self *CourseService) GetCourse(course models.Course) res.Response {
	if err := configs.DB.Where("id = ?", course.ID).Find(&course).Error; err != nil {
		return res.CreateResponseErr(status.ServerError, "internal server error", err)
	}

	return res.CreateResponse(status.Ok, "sucessfully fetch course", course)
}

func (self *CourseService) UpdateCourse(course models.Course) res.Response {
	if _, err := govalidator.ValidateStruct(course); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	if configs.DB.Model(&course).Where("id = ?", course.ID).Updates(&course).RowsAffected == 0 {
		return res.CreateResponseErr(status.ServerError, "failed to update data", fmt.Errorf("data value doesnt change probably"))
	}

	return res.CreateResponse(status.Ok, "course data updated", gin.H {
		"record_data": course,
	})
}

func (self *CourseService) DeleteService(course models.Course) res.Response {
	if configs.DB.Unscoped().Where("id = ?", course.ID).Delete(&course).RowsAffected == 0 {
		return res.CreateResponseErr(status.ServerError, "failed to delete data", fmt.Errorf("data probably didnt exist"))
	}

	return res.CreateResponse(status.Ok, "successfully delete course", nil)
}