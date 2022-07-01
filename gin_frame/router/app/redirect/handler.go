package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com")

}
