package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteAdmin() {
	adminSvc := services.NewAdminService(r.Db)
	adminCtr := controllers.NewAdminController(adminSvc)

	admin := r.Router.Group("/admin")

	admin.POST(
		"/add/student/:studentId/class/:classId", 
		r.AuthMdlwr.RequireAdmin,
		adminCtr.AddStudent,
	)
	admin.DELETE(
		"/remove/student/:studentId/class/:classId",
		r.AuthMdlwr.RequireAdmin,
		adminCtr.RemoveStudent,
	)
}