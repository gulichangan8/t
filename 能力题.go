package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	var times = make([]int, a)
	var ages = make([]int, a)
	var quality []int
	for i := 0; i < a; i++ {
		fmt.Scan(&times[i], &ages[i])
		for j := 0; j < ages[i]; j++ {
			var q int
			fmt.Scan(&q)
			quality = append(quality, q)
		}
	}
	var qs []int
	for i := 0; i < ages[0]; i++ {
		qs = append(qs, quality[i])
	}
	c := repeat(qs)
	fmt.Println(ages[0] - c)
	for i := 1; i < a; i++ {
		var age int
		var begin int
		for j := i - 1; j >= 0; j-- {
			if times[i]-86400 <= times[j] {
				age += ages[j]
				begin = j
			}
		}
		age = age + times[i]
		b := ages[begin]
		var qu = make([]int, age+1)
		for n := b; n < b+age; n++ {
			qu = append(qu, quality[n])
		}
		c := repeat(qu)
		fmt.Println(age - c)
	}
}

func repeat(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for n := i + 1; n < len(arr); n++ {
			if arr[i] == arr[n] {
				count++
			}
		}
	}
	return count
}
