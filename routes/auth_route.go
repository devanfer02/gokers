package router

import (
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

func InitRouteAuth(router *gin.Engine, db *configs.Database) {
	authController := controllers.AuthController{
		Router: router,
		Service: services.AuthService{
			Db: db,
		},
	}

	authStudent := authController.Router.Group("/auth/student")

	authStudent.POST("/register", authController.RegisterStudent)
	authStudent.POST("/login", authController.LoginStudent)
	authStudent.POST("/logout", authController.LogoutStudent)

	authLecturer := router.Group("/auth/lecturer")
	authLecturer.POST("/register", authController.RegisterLecturer)
}