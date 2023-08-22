package controllers

import (
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

type KrsController struct {
	Service services.KrsService
}

func (krsCtr *KrsController) GetKrs(ctx *gin.Context) {

}

func (krsCtr *KrsController) AddClass(ctx *gin.Context) {

}

func (krsCtr *KrsController) RemoveClass(ctx *gin.Context) {

}