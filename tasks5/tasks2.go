package main

import "fmt"

func main() {
	var Bil Person
	Bil.Name = "Bil"
	Bil.age = 10
	fmt.Println(Bil)
	Bil.PrintAge()
	Bil.PrintName()
}

type Person struct {
	Name string
	age  int
}

func (p *Person) PrintAge() {
	fmt.Println(p.age)
}
func (n *Person) PrintName() {
	fmt.Println(n.Name)
}
