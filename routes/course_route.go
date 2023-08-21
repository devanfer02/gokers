package router

import (
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/middlewares"
	"github.com/devanfer02/gokers/services"
)

func (r *Router) InitRouteCourse() {
	courseController := controllers.CourseController{
		Service: services.CourseService{
			Db: r.Db,
		},
	}

	middleware := middlewares.AuthMiddleware{
		Db: r.Db,
	}

	course := r.Router.Group("/course")

	course.GET("/",courseController.GetCourses)
	course.GET("/:id", courseController.GetCourse)
	course.POST("/", middleware.RequireAdmin, courseController.RegisterCourse)
	course.PATCH("/:id", middleware.RequireAdmin, courseController.UpdateCourse)
	course.DELETE("/:id", middleware.RequireAdmin, courseController.DeleteCourse)
}