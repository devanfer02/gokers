package router

import (
	
)

func (r *Router) InitRouteStudent() {
	Student := r.Router.Group("/student")

	Student.GET("/krs")
}