package controllers

import (
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Route 		*gin.Engine 
	Service 	services.AuthService
}

func (auth *AuthController) RegisterStudent(ctx *gin.Context) {
	var student models.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		res.SendResponse(ctx,  res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return 
	}

	response := auth.Service.RegisterStudent(student)

	res.SendResponse(ctx, response)
}

func (auth *AuthController) LoginStudent(ctx *gin.Context) {

}

func (auth *AuthController) LogoutStudent(ctx *gin.Context) {

}

func (auth *AuthController) LoginAdmin(ctx *gin.Context) {

}

func (auth *AuthController) LogoutAdmin(ctx *gin.Context) {

}