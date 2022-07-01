package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "helloHandler  shop",
	})
}
func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "goodsHandler shop",
	})
}
func Routers(e *gin.Engine) {

	e.GET("/hello", helloHandler)
	e.GET("/goods", goodsHandler)

}
