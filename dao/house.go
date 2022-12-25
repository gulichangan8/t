package dao

import (
	"strings"
	"t/model"
)

func AddName(name string) bool {
	_, err := dB.Exec("insert into house (name, good_num, good) value (?,?,?)",
		name, "", "")
	if err != nil {
		return false
	}
	return true
}

func Grounding(name string, ok string) bool {
	_, err := dB.Exec("insert into good (name,state) value (?,?)",
		name, ok)
	if err != nil {
		return false
	}
	return true
}

func Inventory(goodName, houseName string) bool {
	row := dB.QueryRow("select * from house where name=?", houseName)
	var h model.House
	err := row.Scan(&h.Name, &h.GoodNum, &h.Good)
	if err != nil {
		return false
	}
	goods := strings.Split(h.Good, goodName)
	good1 := goods[0]
	good2 := goods[1]
	good := good1 + good2
	_, err = dB.Exec("update house set good_num=?,good=? where name=?", h.GoodNum-1, good, goodName)
	if err != nil {
		return false
	}
	return true
}

func Trans(name, houseNameFrom, houseNameTo string) bool {
	row := dB.QueryRow("select * from house where name=?", houseNameFrom)
	var h model.House
	err := row.Scan(&h.Name, &h.GoodNum, &h.Good)
	if err != nil {
		return false
	}
	goods := strings.Split(h.Good, name)
	good1 := goods[0]
	good2 := goods[1]
	good := good1 + good2
	_, err = dB.Exec("update house set good_num=?,good=? where name=?", h.GoodNum-1, good, houseNameFrom)
	if err != nil {
		return false
	}
	row = dB.QueryRow("select * from house where name=?", houseNameTo)
	err = row.Scan(&h.Name, &h.GoodNum, &h.Good)
	if err != nil {
		return false
	}
	good = h.Good + name
	_, err = dB.Exec("update house set good_num=?,good=? where name=?", h.GoodNum+1, good, houseNameTo)
	if err != nil {
		return false
	}
	return true
}
