package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouteStudent(router *gin.Engine) {
	Student := router.Group("/student")

	Student.GET("/krs")
}