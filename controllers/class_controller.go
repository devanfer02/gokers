package controllers

import (
	"github.com/devanfer02/gokers/models"
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

type ClassController struct {
	Service services.ClassService
}

func (classCtr *ClassController) RegisterClass (ctx *gin.Context) {
	courseId, err := helpers.GetParamID(ctx)	

	if err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad param request", err))
		return
	}

	var class models.Class
	class.CourseID = courseId

	if err := ctx.ShouldBindJSON(&class); err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return
	}

	response := classCtr.Service.RegisterClass(class);

	res.SendResponse(ctx, response)
}