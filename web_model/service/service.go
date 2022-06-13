package service

import (
	"golang/web_model/model"
	"golang/web_model/pkg/errno"

	"golang/web_model/handler"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var r model.User
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.User{
		UserName: r.UserName,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, u)
}

func SelectUser(c *gin.Context) {
	name := c.Query("user_name")
	if name == "" {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	var user model.User

	if err := user.Validate(); err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SelectUser(c, nil, user)
}
