package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func bindUrlHandler(c *gin.Context) {
	//声明接收的变量
	var login Login

	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if login.User != "root" || login.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
