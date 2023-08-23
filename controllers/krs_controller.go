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

	response := krsCtr.Service.GetKrs(krs, studentId)

	res.SendResponse(ctx, response)
}

func (krsCtr *KrsController) AddClass(ctx *gin.Context) {

}

func (krsCtr *KrsController) RemoveClass(ctx *gin.Context) {

}