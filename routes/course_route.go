package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteCourse() {
	courseSvc := services.NewCourseService(r.Db)
	courseCtr := controllers.NewCourseController(courseSvc)

	course := r.Router.Group("/course")

	course.GET("/", r.AuthMdlwr.RequireAdmin, courseCtr.GetCourses)
	course.GET("/:id", r.AuthMdlwr.RequireAdmin, courseCtr.GetCourse)
	course.POST("/", r.AuthMdlwr.RequireAdmin, courseCtr.RegisterCourse)
	course.PATCH("/:id", r.AuthMdlwr.RequireAdmin, courseCtr.UpdateCourse)
	course.DELETE("/:id", r.AuthMdlwr.RequireAdmin, courseCtr.DeleteCourse)
}