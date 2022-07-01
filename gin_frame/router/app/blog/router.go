package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {

	e.GET("/post", postHandler)

}
func postHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "postHandler www.topgoer.com",
	})
}
