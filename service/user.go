package service

import "t/dao"

func CheckLogin(username string, password string) bool {
	U := dao.TakeUserDate()
	ok := false
	for _, date := range U {
		if date.Username == username && date.Password == password {
			ok = true
		} else {
			continue
		}
	}
	if ok {
		return true
	} else {
		return false
	}
}
