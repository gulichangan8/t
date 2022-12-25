package service

import "t/dao"

func AddHouse(name string) bool {
	ok := dao.AddName(name)
	return ok
}

func Grounding(name string, ok string) bool {
	OK := dao.Grounding(name, ok)
	return OK
}

func Inventory(goodName, houseName string) bool {
	ok := dao.Inventory(goodName, houseName)
	return ok
}

func Trans(name, houseNameFrom, houseNameTo string) bool {
	ok := dao.Trans(name, houseNameFrom, houseNameTo)
	return ok
}
