package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var channel = make(chan bool)
	wg.Add(3)
	go Work("goroutine1", channel)
	go Work("goroutine2", channel)
	go Work("goroutine3", channel)
	wg.Wait()
	fmt.Println("successful")
}

func Work(workName string, channel chan bool) {
	var c = true
	channel <- c
	close(channel)
	time.Sleep(time.Second) // 模拟逻辑处理
	fmt.Println(workName)
	wg.Done()
}

//goroutine1
//goroutine2
//goroutine3
//successful
