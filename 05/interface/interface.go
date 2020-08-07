package main

import "fmt"

func main() {
	printType(3)
	printType("интерфейсы - это легко")
	printType([]string{"слайсы", "тоже"})
}

func printType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("тип аргумента int")
	case string:
		fmt.Println("тип аргумента string")
	default:
		fmt.Println("значение аргумента не int и не string")
	}
}
