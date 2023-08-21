package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteAuth() {
	authController := controllers.AuthController{
		Service: services.AuthService{
			Db: r.Db,
		},
	}

	authStudent := r.Router.Group("/auth/student")

	authStudent.POST("/register", authController.RegisterStudent)
	authStudent.POST("/login", authController.LoginStudent)
	authStudent.POST("/logout", authController.LogoutStudent)

	authLecturer := r.Router.Group("/auth/lecturer")
	authLecturer.POST("/register", authController.RegisterLecturer)

	authAdmin := r.Router.Group("/auth/admin")
	authAdmin.POST("/register", authController.RegisterAdmin)
	authAdmin.POST("/login", authController.LoginAdmin)
	authAdmin.POST("/logout", authController.LogoutAdmin)
}