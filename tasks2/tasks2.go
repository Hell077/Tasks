package main

import (
	"fmt"
)

func main() {
	fmt.Println(week(1))
	fmt.Println(week(2))
	fmt.Println(week(3))
	fmt.Println(week(4))
	fmt.Println(week(5))
	fmt.Println(week(6))
	fmt.Println(week(7))

}
func week(num int) string {
	switch num {
	case 1:
		return "понедельник"

	case 2:
		return "вторник"

	case 3:
		return "среда"

	case 4:
		return "четверг"

	case 5:
		return "пятница"

	case 6:
		return "суббота"

	case 7:
		return "воскресение"
	default:
		return "ошибка"
	}

}
