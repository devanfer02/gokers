package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteAuth() {
	authSvc := services.NewAuthService(r.Db)
	authCtr := controllers.NewAuthController(authSvc)

	authStudent := r.Router.Group("/auth/student")

	authStudent.POST("/register", authCtr.RegisterStudent)
	authStudent.POST("/login", authCtr.LoginStudent)
	authStudent.POST("/logout", authCtr.LogoutStudent)

	authLecturer := r.Router.Group("/auth/lecturer")
	authLecturer.POST("/register", authCtr.RegisterLecturer)

	authAdmin := r.Router.Group("/auth/admin")
	authAdmin.POST("/register", authCtr.RegisterAdmin)
	authAdmin.POST("/login", authCtr.LoginAdmin)
	authAdmin.POST("/logout", authCtr.LogoutAdmin)
}