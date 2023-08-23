package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/middlewares"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteKrs() {
	krsSvc := services.NewKrsService(r.Db)
	krsCtr := controllers.NewKrsController(krsSvc)

	middleware := middlewares.AuthMiddleware{
		Db: r.Db, 
	}

	krs := r.Router.Group("/krs")

	krs.GET("/", middleware.RequireAuth, krsCtr.GetKrs)
	krs.POST("/:classId", middleware.RequireAuth, krsCtr.AddClass)
	krs.DELETE("/:classId", middleware.RequireAuth, krsCtr.RemoveClass)
}