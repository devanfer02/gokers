package controllers

import (
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"

	//"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	Service *services.AdminService
}

func NewAdminController(adminSvc *services.AdminService) *AdminController {
	return &AdminController{adminSvc}
}

func (admCtr *AdminController) AddStudent(ctx *gin.Context) {
	studentId, err := helpers.GetParamUUID(ctx, "studentId")

	if err != nil {
		res.SendResponse(
			ctx, 
			res.CreateResponseErr(status.BadRequest, "bad param request", err),
		)
	}

	classId, err := helpers.GetParamUUID(ctx, "classId")

	if err != nil {
		res.SendResponse(
			ctx, 
			res.CreateResponseErr(status.BadRequest, "bad param request", err),
		)
	}

	var student models.Student 
	var class models.Class 

	student.ID = studentId 
	class.ID = classId 

	response := admCtr.Service.AddStudent(&student, &class)

	res.SendResponse(ctx, response)
}

func (admCtr *AdminController) RemoveStudent(ctx *gin.Context) {
	studentId, err := helpers.GetParamUUID(ctx, "studentId")

	if err != nil {
		res.SendResponse(
			ctx, 
			res.CreateResponseErr(status.BadRequest, "bad param request", err),
		)
	}

	classId, err := helpers.GetParamUUID(ctx, "classId")

	if err != nil {
		res.SendResponse(
			ctx, 
			res.CreateResponseErr(status.BadRequest, "bad param request", err),
		)
	}

	var student models.Student 
	var class models.Class 

	student.ID = studentId 
	class.ID = classId 

	response := admCtr.Service.RemoveStudent(&student, &class)

	res.SendResponse(ctx, response)
}