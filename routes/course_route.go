package router

import (
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

func InitRouteCourse(router *gin.Engine, db *configs.Database) {
	courseController := controllers.CourseController{
		Router: router,
		Service: services.CourseService{
			Db: db,
		},
	}

	course := courseController.Router.Group("/course")

	//TODO: Make admin only
	course.GET("/", courseController.GetCourses)
	course.GET("/:id", courseController.GetCourse)
	course.POST("/", courseController.RegisterCourse)
	course.PATCH("/:id", courseController.UpdateCourse)
	course.DELETE("/:id", courseController.DeleteCourse)
}