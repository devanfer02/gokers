package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteClass() {
	classController := controllers.ClassController{
		Service: services.ClassService{
			Db: r.Db,
		},
	}

	class := r.Router.Group("/class")

	class.GET("/", r.AuthMdlwr.RequireAuth, classController.GetClasses)
	class.GET("/:id", r.AuthMdlwr.RequireAuth, classController.GetClass)
	class.POST("/", r.AuthMdlwr.RequireAdmin, classController.RegisterClass)
	class.PATCH("/:id", r.AuthMdlwr.RequireAdmin, classController.UpdateClass)
	class.DELETE("/:id", r.AuthMdlwr.RequireAdmin, classController.DeleteClass)
}