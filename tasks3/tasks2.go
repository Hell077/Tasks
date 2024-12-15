package main

import "fmt"

func main() {
	arr := []int{16671, 352, 3432, 4, 55454, 29322312, 88}
	var maxVal int
	for _, v := range arr {
		if maxVal < v {
			maxVal = v
		}
	}
	fmt.Println(maxVal)
}
