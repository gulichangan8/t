package main

import "fmt"

// defer将"1"的打印延后，先执行下面的"2"的打印，打印完"2"return，此时打印3的函数无法执行再执行1
// 所以没有输出3
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()

	if a {
		fmt.Println("2")
		//return
	}

	defer func() {
		fmt.Println("3")
	}()
}
