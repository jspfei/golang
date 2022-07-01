package main

import (
	"fmt"
	"golang/gin_frame/router/app/async"
	"golang/gin_frame/router/app/blog"
	"golang/gin_frame/router/app/htmltemplate"
	"golang/gin_frame/router/app/login"
	"golang/gin_frame/router/app/redirect"
	"golang/gin_frame/router/app/response"
	"golang/gin_frame/router/app/shop"
	"golang/gin_frame/router/routers"
)

func main() {
	routers.Include(response.Routers, redirect.Routers, async.Routers, shop.Routers, blog.Routers, login.Routers, htmltemplate.Routers)

	r := routers.Init()

	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err :%v\n", err)
	}
}
