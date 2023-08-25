package res

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func CreateResponse(code int, message string, data interface{}) Response {
	return Response{
		Status:  "success",
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func CreateResponseErr(code int, message string, err error) Response {
	var errdata interface{}

	if err != nil {
		errdata = err.Error()
	} else {
		errdata = nil 
	}

	return Response{
		Status:  "error",
		Code:    code,
		Message: message,
		Data: errdata, 
	}
}

func SendResponse(ctx *gin.Context, response Response) {
	ctx.JSON(response.Code, response)
}

func SendStatusOnly(ctx *gin.Context, code int) {
	ctx.AbortWithStatus(code)
}
