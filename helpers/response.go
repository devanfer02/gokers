package helpers

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Status 	string		`json:"status"`
	Code	int			`json:"code"`
	Message string		`json:"message"`
	Data 	interface{}	`json:"data,omitempty"`
}

func Response(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, response {
		Status: "success",
		Code: code, 
		Message: message, 
		Data: data,
	})
}

func ResponseErr(ctx *gin.Context, code int, message string, err error) {
	ctx.AbortWithStatusJSON(code, response {
		Status: "error", 
		Code: code, 
		Message: message, 
		Data: err.Error(),
	})
}
