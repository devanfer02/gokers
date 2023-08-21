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

	if token == "" {
		res.SendResponse(ctx, response)
		return
	}

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

func (authCtr *AuthController) RegisterAdmin(ctx *gin.Context) {
	var admin models.Admin

	if err := ctx.ShouldBindJSON(&admin); err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return 
	}

	response := authCtr.Service.RegisterAdmin(admin);

	res.SendResponse(ctx, response)
}

func (authCtr *AuthController) LoginAdmin(ctx *gin.Context) {
	var admin models.Admin
	
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return
	}

	response, token := authCtr.Service.LoginAdmin(admin)

	if token == "" {
		res.SendResponse(ctx, response)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("AdminAuthorization", token, 3600 * 6, "", "", true, true)

	res.SendResponse(ctx, response)
}

func (autCtr *AuthController) LogoutAdmin(ctx *gin.Context) {
	ctx.Set("adm", nil) 
	ctx.SetCookie("AdminAuthorization", "", -1, "", "", true, true)

	res.SendResponse(ctx, res.CreateResponse(status.Ok, "admin successfully logout", nil))
}