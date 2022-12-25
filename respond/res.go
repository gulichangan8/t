package respond

import "github.com/gin-gonic/gin"

func LoginTrue(c *gin.Context) {
	c.String(200, "登陆成功")
}

func LoginErr(c *gin.Context) {
	c.String(200, "登陆失败")
}

func AddHouseTrue(c *gin.Context) {
	c.String(200, "新增仓库成功")
}

func AddHouseErr(c *gin.Context) {
	c.String(200, "新增仓库失败")
}

func ChangeGoodTrue(c *gin.Context) {
	c.String(200, "货物已上/下架")
}

func ChangeGoodErr(c *gin.Context) {
	c.String(200, "货物上/下架失败")
}

func InventoryTrue(c *gin.Context) {
	c.String(200, "货物出/入库成功")
}

func InventoryErr(c *gin.Context) {
	c.String(200, "货物出/入库失败")
}

func TransTrue(c *gin.Context) {
	c.String(200, "货物调库成功")
}

func TransErr(c *gin.Context) {
	c.String(200, "货物调库失败")
}
