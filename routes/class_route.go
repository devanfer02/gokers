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
	class.POST("/", r.AuthMdlwr.RequireAdmin, classCtr.RegisterClass)
	class.PATCH("/:id", r.AuthMdlwr.RequireAdmin, classCtr.UpdateClass)
	class.DELETE("/:id", r.AuthMdlwr.RequireAdmin, classCtr.DeleteClass)
}