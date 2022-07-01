package htmltemplate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.LoadHTMLGlob("app/htmltemplate/*")
	//	e.Static("/assets", "./assets")
	e.GET("/index", htmlTemplateHandler)
}

func htmlTemplateHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "测试", "ce": "122341412"})
}
