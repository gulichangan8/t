package api

import (
	"github.com/gin-gonic/gin"
	"t/respond"
	"t/service"
)

func AddHouse(c *gin.Context) {
	name := c.PostForm("house_name")
	ok := service.AddHouse(name)
	if ok {
		respond.AddHouseTrue(c)
	} else {
		respond.AddHouseErr(c)
	}
}

func ChangeGood(c *gin.Context) {
	name := c.PostForm("good_name")
	ok := c.PostForm("is_grounding")
	OK := service.Grounding(name, ok)
	if OK {
		respond.ChangeGoodTrue(c)
	} else {
		respond.ChangeGoodErr(c)
	}
}

func Inventory(c *gin.Context) {
	goodName := c.PostForm("good_name")
	houseName := c.PostForm("house_name")
	ok := service.Inventory(goodName, houseName)
	if ok {
		respond.InventoryTrue(c)
	} else {
		respond.InventoryErr(c)
	}
}

func Trans(c *gin.Context) {
	name := c.PostForm("good_name")
	houseNameFrom := c.PostForm("house_name_from")
	houseNameTo := c.PostForm("house_name_to")
	ok := service.Trans(name, houseNameFrom, houseNameTo)
	if ok {
		respond.TransTrue(c)
	} else {
		respond.TransErr(c)
	}
}
