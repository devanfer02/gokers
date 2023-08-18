package main

import (
	"os"

	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	configs.ConnectToDB()
	configs.MigrateDB()

	app := gin.Default()

	router.InitRouteAuth(app)

	app.Run(os.Getenv("APP_PORT"))
}