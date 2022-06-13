package router

import (
	"golang/web_model/router/middleware"
	"golang/web_model/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	middlewares := []gin.HandlerFunc{}
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middlewares...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	router := g.Group("/")
	{
		router.POST("/addUser", service.AddUser)
		router.POST("/selectUser", service.SelectUser)
	}
}
