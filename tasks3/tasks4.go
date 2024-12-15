package main

import (
	"fmt"
	"sort"
)

func main() {
	some := []int{1, 23, 24, 3, 5342, 65, 547, 45, 7, 54, 7, 545, 423, 345, 2563, 34, 7, 56845}
	fmt.Println(some)
	sort.Ints(some)
	fmt.Println(some)
}
