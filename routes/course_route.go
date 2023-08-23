package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteCourse() {
	courseController := controllers.CourseController{
		Service: services.CourseService{
			Db: r.Db,
		},
	}

	course := r.Router.Group("/course")

	course.GET("/", r.AuthMdlwr.RequireAdmin, courseController.GetCourses)
	course.GET("/:id", r.AuthMdlwr.RequireAdmin, courseController.GetCourse)
	course.POST("/", r.AuthMdlwr.RequireAdmin, courseController.RegisterCourse)
	course.PATCH("/:id", r.AuthMdlwr.RequireAdmin, courseController.UpdateCourse)
	course.DELETE("/:id", r.AuthMdlwr.RequireAdmin, courseController.DeleteCourse)
}