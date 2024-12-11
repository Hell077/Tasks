package main

import "fmt"

func main() {
	res := sum(100, 2)
	fmt.Println(res)
	fmt.Println(sum(100, 2))
}

func sum(first, second int) int {
	return first + second

}
