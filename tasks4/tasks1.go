package main

import (
	"fmt"
)

func main() {
	fmt.Println(Factorial(10))
}

func Factorial(num int64) int64 {
	if num == 0 || num == 1 {
		return 1
	}
	return num * Factorial(num-1)
}
