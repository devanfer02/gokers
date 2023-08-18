package res

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func CreateResponseOk(code int, message string, data interface{}) Response {
	return Response{
		Status:  "success",
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func CreateResponseErr(code int, message string, err error) Response {
	return Response{
		Status:  "error",
		Code:    code,
		Message: message,
		Data:    err.Error(),
	}
}

func SendResponse(ctx *gin.Context, response Response) {
	ctx.JSON(response.Code, response)
}
