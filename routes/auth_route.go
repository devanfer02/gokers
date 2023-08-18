package router

import (
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func InitRouteAuth(router *gin.Engine) {
	authController := controllers.AuthController{
		Route: router,
		Service: services.AuthService{},
	}

	authStudent := router.Group("/auth/student")

	authStudent.POST("/register", authController.RegisterStudent)
// 	authStudent.POST("/login")
// 	authStudent.POST("/logout")

// 	authAdmin := router.Group("/auth/admin")

// 	authAdmin.POST("/register")
// 	authAdmin.POST("/login")
// 	authAdmin.POST("/logout")
}