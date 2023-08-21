package controllers

import (
	"net/http"

	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Router 		*gin.Engine 
	Service 	services.AuthService
}

func (authCtr *AuthController) RegisterStudent(ctx *gin.Context) {
	var student models.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return 
	}

	response := authCtr.Service.RegisterStudent(student)

	res.SendResponse(ctx, response)
}

func (authCtr *AuthController) LoginStudent(ctx *gin.Context) {
	var studentAuth models.StudentAuth

	if err := ctx.ShouldBindJSON(&studentAuth); err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return
	}

	response, token := authCtr.Service.LoginStudent(studentAuth)

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600 * 12, "", "", true, true)

	res.SendResponse(ctx, response)
}

func (authCtr *AuthController) LogoutStudent(ctx *gin.Context) {
	ctx.Set("std", nil)
	ctx.SetCookie("Authorization", "", -1, "", "", true, true)

	res.SendResponse(ctx, res.CreateResponse(status.Ok, "student successfully logout", nil))
}

func (authCtr *AuthController) RegisterLecturer(ctx *gin.Context) {
	var lecturer models.Lecturer

	if err := ctx.ShouldBindJSON(&lecturer); err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return 
	}

	response := authCtr.Service.RegisterLecturer(lecturer)

	res.SendResponse(ctx, response)
}