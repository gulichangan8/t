package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	Username string
	Nickname string
	Sex      uint8
	Birthday time.Time
}

// 首字母需要大写才能被json包访问到
func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}
