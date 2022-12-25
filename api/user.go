package api

import (
	"github.com/gin-gonic/gin"
	"t/respond"
	"t/service"
)

func Login(c *gin.Context) {
	u := c.PostForm("username")
	p := c.PostForm("password")
	ok := service.CheckLogin(u, p)
	if ok {
		respond.LoginTrue(c)
	} else {
		respond.LoginErr(c)
	}
}
