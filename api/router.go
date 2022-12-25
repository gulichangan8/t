package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/login", Login)
	}

	house := r.Group("/house")
	{
		house.POST("/add", AddHouse)
		house.POST("/good_change", ChangeGood)
		house.GET("/goods", Inventory)
		house.POST("/trans", Trans)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
