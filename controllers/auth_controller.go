package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/gokers/helpers"
)

func RegisterStudent(ctx *gin.Context) {
	helpers.Response(ctx, 200, "hello world!", "ini data")	
}

func LoginStudent(ctx *gin.Context) {

}

func LogoutStudent(ctx *gin.Context) {

}

func LoginAdmin(ctx *gin.Context) {

}

func LogoutAdmin(ctx *gin.Context) {

}