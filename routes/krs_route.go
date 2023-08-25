package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteKrs() {
	krsSvc := services.NewKrsService(r.Db)
	krsCtr := controllers.NewKrsController(krsSvc)

	krs := r.Router.Group("/krs")

	krs.GET("/", r.AuthMdlwr.RequireAuth, krsCtr.GetKrs)
	krs.POST("/:classId", r.AuthMdlwr.RequireAuth, krsCtr.AddClass)
	krs.DELETE("/:classId", r.AuthMdlwr.RequireAuth, krsCtr.RemoveClass)
}