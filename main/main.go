package main

import (
	"t/api"
	"t/dao"
)

func main() {
	dao.InitMySql()
	api.InitEngine()
}
