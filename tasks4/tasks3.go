package main

import "fmt"

func main() {
	str := "Hello World"
	fmt.Println(str)
	fmt.Println(Reverse(str))
}
func Reverse(str string) string {
	runes := []rune(str)
	if len(str) == 0 || len(str) == 1 {
		return str
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
