package main

import (
	"os"

	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/middlewares"
	"github.com/devanfer02/gokers/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	app 		:= gin.Default()
	db 			:= &configs.Database{}
	authMdlwr 	:= &middlewares.AuthMiddleware{Db: db}
	router := router.Router{
		Router: app, 
		Db: db,
		AuthMdlwr: authMdlwr,
	}

	db.ConnectToDB()
	db.MigrateDB()

	configs.InitMajors()
	
	app.Use(middlewares.InterceptApi)

	router.InitRouteAuth()
	router.InitRouteCourse()

	app.Run(os.Getenv("APP_PORT"))
}