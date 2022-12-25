package main

import (
	"fmt"
	"math"
	"sync"
)

var w sync.WaitGroup

func main() {
	w.Add(6)
	go func() {
		for i := 0; i < 10; i++ {
			if i == 1 {
				fmt.Println(i)
			}
		}
		w.Done()
	}()
	go func() {
		for i := 10; i < 100; i++ {
			if int(math.Pow(float64(i%10), 2)+math.Pow(float64(i/10), 2)) == i {
				fmt.Println(i)
			}
		}
		w.Done()
	}()
	go func() {
		for i := 100; i < 1000; i++ {
			if int(math.Pow(float64((i%100)%10), 3)+math.Pow(float64((i%100)/10), 3)+math.Pow(float64(i/100), 3)) == i {
				fmt.Println(i)
			}
		}
		w.Done()
	}()
	go func() {
		for i := 1000; i < 10000; i++ {
			if int(math.Pow(float64(((i%1000)%100)%10), 4)+math.Pow(float64(((i%1000)%100)/10), 4)+
				math.Pow(float64((i%1000)/100), 4)+math.Pow(float64(i/1000), 4)) == i {
				fmt.Println(i)
			}
		}
		w.Done()
	}()
	go func() {
		for i := 10000; i < 100000; i++ {
			if int(math.Pow(float64((((i%10000)%1000)%100)%10), 5)+math.Pow(float64((((i%10000)%1000)%100)/10), 5)+
				math.Pow(float64(((i%10000)%1000)/100), 5)+math.Pow(float64((i%10000)/1000), 5)+
				math.Pow(float64(i/10000), 5)) == i {
				fmt.Println(i)
			}
		}
		w.Done()
	}()
	go func() {
		for i := 100000; i < 1000000; i++ {
			if int(math.Pow(float64(((((i%100000)%10000)%1000)%100)%10), 6)+math.Pow(float64(((((i%100000)%10000)%1000)%100)/10), 6)+
				math.Pow(float64((((i%100000)%10000)%1000)/100), 6)+math.Pow(float64(((i%100000)%10000)/1000), 6)+
				math.Pow(float64((i%100000)/10000), 6)+math.Pow(float64(i/100000), 6)) == i {
				fmt.Println(i)
			}
		}
		w.Done()
	}()
	w.Wait()
}
