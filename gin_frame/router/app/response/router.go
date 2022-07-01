package response

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/someJSON", jsonHandler)
	e.GET("/someStruct", someStructHandler)
	e.GET("/someXML", someXmlHandler)
	e.GET("/someYAML", someYamlHandler)
	e.GET("/someProtoBuf", someProtoBufHandler)
}
