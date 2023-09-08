package controllers

import (
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/models"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

type LecturerController struct {
	Service *services.LecturerService
}

func NewLecturerController(lctSvc *services.LecturerService) *LecturerController {
	return &LecturerController{Service: lctSvc}
}

func (lctCtr *LecturerController) GetLecturers(ctx *gin.Context) {
	var lecturers []models.Lecturer

	response := lctCtr.Service.GetLecturers(&lecturers) 

	res.SendResponse(ctx, response)
}