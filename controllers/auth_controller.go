package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/services"
)

type AuthController struct {
	Route 		*gin.Engine 
	Service 	services.AuthService
}

func (auth *AuthController) RegisterStudent(ctx *gin.Context) {
	helpers.Response(ctx, 200, "hello world!", "ini data")	
}

func (auth *AuthController) LoginStudent(ctx *gin.Context) {

}

func (auth *AuthController) LogoutStudent(ctx *gin.Context) {

}

func (auth *AuthController) LoginAdmin(ctx *gin.Context) {

}

func (auth *AuthController) LogoutAdmin(ctx *gin.Context) {

}