package router

import (
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/middlewares"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Router		*gin.Engine 
	Db 			*configs.Database
	AuthMdlwr	*middlewares.AuthMiddleware
}

func NewRouter(router *gin.Engine, db *configs.Database, authMdlwr *middlewares.AuthMiddleware) *Router {
	return &Router{
		Router: router, 
		Db: db,
		AuthMdlwr: authMdlwr,
	}
}