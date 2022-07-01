package middle

//中间件

import (
	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {

}

func Routers(e *gin.Engine) {
	r.Use(MiddleWare())
}
