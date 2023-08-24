package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteClass() {
	classSvc := services.NewClassService(r.Db)
	classCtr := controllers.NewClassController(classSvc)

	class := r.Router.Group("/class")

	class.GET("/", r.AuthMdlwr.RequireAuth, classCtr.GetClasses)
	class.GET("/:id", r.AuthMdlwr.RequireAuth, classCtr.GetClass)
	class.GET("/admin", r.AuthMdlwr.RequireAdmin, classCtr.GetClasses)
	class.GET("/admin/:id", r.AuthMdlwr.RequireAdmin, classCtr.GetClass)
	class.POST("/register", r.AuthMdlwr.RequireAdmin, classCtr.RegisterClass)
	class.PATCH("/update/:id", r.AuthMdlwr.RequireAdmin, classCtr.UpdateClass)
	class.DELETE("/delete/:id", r.AuthMdlwr.RequireAdmin, classCtr.DeleteClass)
}