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

	db := configs.Database{}

	db.ConnectToDB()
	db.MigrateDB()

	configs.InitMajors()

	app := gin.Default()
	app.Use(middlewares.InterceptApi)

	router.InitRouteAuth(app, &db)

	app.Run(os.Getenv("APP_PORT"))
}