package controllers

import (
	//"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	//"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	Service *services.StudentService
}

func (stdCtr *StudentController) UpdateStudent(ctx *gin.Context) {
	var student models.Student

	response := stdCtr.Service.UpdateStudent(&student)

	res.SendResponse(ctx, response)
}