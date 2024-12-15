package main

import "fmt"

func main() {
	fmt.Println(check(75))
}
func check(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
