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

	//TODO: Make admin only
	course.GET("/", courseController.GetCourses)
	course.GET("/:id", courseController.GetCourse)
	course.POST("/", courseController.RegisterCourse)
	course.PATCH("/:id", courseController.UpdateCourse)
	course.DELETE("/:id", courseController.DeleteCourse)
}