package router

import (
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func InitRouteCourse(router *gin.Engine) {
	courseController := controllers.CourseController{
		Router: router,
		Service: services.CourseService{},
	}

	course := courseController.Router.Group("/course")

	course.GET("/", courseController.GetCourses)
	course.GET("/:id", courseController.GetCourse)
	course.POST("/", courseController.RegisterCourse)
	course.PATCH("/:id", courseController.UpdateCourse)
	course.DELETE("/:id", courseController.DeleteCourse)
}