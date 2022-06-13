package wrapper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// gin web出参封装
func wrap(context *gin.Context, code int, message string, data interface{}) {
	context.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Custom(context *gin.Context, code int, message string) {
	wrap(context, code, message, nil)
}

// Default 默认出参
func Default(context *gin.Context) {
	wrap(context, http.StatusOK, "success", nil)
}

// Success 执行成功出参
func Success(context *gin.Context, data interface{}) {
	wrap(context, http.StatusOK, "success", data)
}

// Error 执行异常出参
func Error(context *gin.Context, message string) {
	wrap(context, http.StatusInternalServerError, message, nil)
}
