package main

import "fmt"

func main() {
	var channel = make(chan bool)
	var c = true
	go func() {
		channel <- c
		fmt.Println("出现")
	}()
	b := <-channel
	fmt.Println(b)
}
