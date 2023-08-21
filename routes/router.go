package router

import (
	"github.com/devanfer02/gokers/configs"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Router	*gin.Engine 
	Db 		*configs.Database
}
