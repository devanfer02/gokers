package controllers

import (
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

type KrsController struct {
	Service *services.KrsService
}

func NewKrsController(krsSvc *services.KrsService) *KrsController {
	return &KrsController{Service: krsSvc}
}

func (krsCtr *KrsController) GetKrs(ctx *gin.Context) {
	var krs []models.KRS
	studentId, ok := helpers.GetStudentID(ctx)

	if !ok {
		res.SendStatusOnly(ctx, status.Unauthorized)
		return 
	}

	response := krsCtr.Service.GetKrs(&krs, studentId)

	res.SendResponse(ctx, response)
}

func (krsCtr *KrsController) AddClass(ctx *gin.Context) {
	var krs models.KRS 

	id, err := helpers.GetParamID(ctx)
	studentId, ok := helpers.GetStudentID(ctx)

	if !ok {
		res.SendStatusOnly(ctx, status.Unauthorized)
	}

	if err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad param request", err))
		return 
	}

	krs.ClassID = id 
	krs.StudentID = studentId

	response := krsCtr.Service.AddClass(&krs)

	res.SendResponse(ctx, response)
}

func (krsCtr *KrsController) RemoveClass(ctx *gin.Context) {
	var krs models.KRS

	id, err := helpers.GetParamID(ctx)
	studentId, ok := helpers.GetStudentID(ctx)

	if !ok {
		res.SendStatusOnly(ctx, status.Unauthorized)
		return
	}

	if err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad param request", err))
		return 
	}

	krs.ClassID = id 
	krs.StudentID = studentId

	response := krsCtr.Service.RemoveClass(&krs)

	res.SendResponse(ctx, response)
}