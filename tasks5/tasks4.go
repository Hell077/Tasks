package main

import "fmt"

func main() {
	CheckImp()
	var obj ExampleInterface = MyStruct{}
	fmt.Println(obj.DoSomething())
}

type ExampleInterface interface {
	DoSomething() string
}

type MyStruct struct{}

func (m MyStruct) DoSomething() string {
	return "Моя структура имплементриует интерфейс"
}

func CheckImp() {
	var _ ExampleInterface = MyStruct{}
	fmt.Println("Моя структура успешно имплиметирует интефрейс")
}
