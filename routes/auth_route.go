package router

import (
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/gokers/controllers"
)

func InitRouteAuth(router *gin.Engine) {
	authStudent := router.Group("/auth/student")

	authStudent.POST("/register", controllers.RegisterStudent)
	authStudent.POST("/login")
	authStudent.POST("/logout")

	authAdmin := router.Group("/auth/admin")

	authAdmin.POST("/register")
	authAdmin.POST("/login")
	authAdmin.POST("/logout")
}