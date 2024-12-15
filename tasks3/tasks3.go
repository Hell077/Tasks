package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}
	mergeSlice := append(arr1, arr2...)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(mergeSlice)
}
