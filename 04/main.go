package main

import "fmt"

func main() {
	ages := map[string]int{
		"Максим": 20,
		"Олег": 24,
		"Саня": 28,
		"Антон": 35,
	}

	var ages2 map[string]int

	fmt.Println(ages)
	fmt.Println(ages2)

	ages2 = ages

	delete(ages, "Антон")
	delete(ages, "Саня")

	fmt.Println(ages)
	fmt.Println(ages2)
}