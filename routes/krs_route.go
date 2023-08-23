package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/middlewares"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteKrs() {
	krsController := controllers.KrsController{
		Service: services.KrsService{
			Db: r.Db, 
		}, 
	}

	middleware := middlewares.AuthMiddleware{
		Db: r.Db, 
	}

	krs := r.Router.Group("/krs")

	krs.GET("/", middleware.RequireAuth, krsController.GetKrs)
	krs.POST("/:classId", middleware.RequireAuth, krsController.AddClass)
	krs.DELETE("/:classId", middleware.RequireAuth, krsController.RemoveClass)
}