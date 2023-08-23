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

func (classCtr *ClassController) GetClasses(ctx *gin.Context) {
	var classes []models.Class
	queries := make([]string, 2)

	queries[0] = ctx.Query("code")
	queries[1] = ctx.Query("major")

	response := classCtr.Service.GetClasses(classes, queries)

	res.SendResponse(ctx, response)
}

func (classCtr *ClassController) GetClass(ctx *gin.Context) {
	var class models.Class

	id, err := helpers.GetParamID(ctx)

	if err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad param request", err))
		return
	}

	class.ID = id

	response := classCtr.Service.GetClass(class)

	res.SendResponse(ctx, response)
}

func (classCtr *ClassController) RegisterClass(ctx *gin.Context) {
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

func (classCtr *ClassController) UpdateClass(ctx *gin.Context) {
	var class models.Class 

	if err := ctx.ShouldBindJSON(&class); err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return
	}

	id, err := helpers.GetParamID(ctx)

	if err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad param request", err))
		return
	}

	class.ID = id 

	response := classCtr.Service.UpdateClass(class)

	res.SendResponse(ctx, response)
}

func (classCtr *ClassController) DeleteClass(ctx *gin.Context) {
	id, err := helpers.GetParamID(ctx)

	if err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad param request", err))
		return
	}

	var class models.Class 
	class.ID = id 

	response := classCtr.Service.DeleteClass(class)

	res.SendResponse(ctx, response)	
}