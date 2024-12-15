package main

import "fmt"

func main() {
	var Bil person
	Bil.Name = "Bil"
	Bil.age = 10
	fmt.Println(Bil)
}

type person struct {
	Name string
	age  int
}
