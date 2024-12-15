package main

import (
	"fmt"
	"math"
)

func main() {
	rect := Rectangle{
		width:  10,
		height: 20,
	}
	circ := Circle{
		radius: 30,
	}
	fmt.Println(circ.Area())
	fmt.Println(circ.Perimeter())
	fmt.Println(rect.Perimeter())
	fmt.Println(rect.Area())
}

// Квадрат структура

type Rectangle struct {
	width  float32
	height float32
}

// Круг структура

type Circle struct {
	radius float32
}

// Интерфейс который принимает одинаковые данные для двух одинаковых методов
// Расчёта площади и периметра круга и квадрата

type Shape interface {
	Area() float32
	Perimeter() float32
}

func (r *Rectangle) Area() float32 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float32 {
	return (r.width + r.height) * 2
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}
