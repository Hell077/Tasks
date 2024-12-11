package main

import "fmt"

func main() {
	res := dev(11, 0)
	fmt.Println(res)
	fmt.Println(dev(11, 100))
}
func dev(first, second int) float64 {
	res := float64(first) / float64(second)

	if second == 0 {
		return 0.0
	}
	return res
}
