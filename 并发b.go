package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var channels [3]chan int

func main() {

	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)
	channels[0] = ch1
	channels[1] = ch2
	channels[2] = ch3

	wg.Add(3)
	go work("goroutine1", 0)
	go work("goroutine2", 1)
	go work("goroutine3", 2)
	ch1 <- 1
	wg.Wait()
	fmt.Println("successful")
}

func work(workName string, index int) {
	var ch = channels[index]
	<-ch
	close(ch)

	time.Sleep(time.Second)
	fmt.Println(workName)
	wg.Done()

	index = index + 1
	if index < len(channels) {
		channels[index] <- index
	}
}
