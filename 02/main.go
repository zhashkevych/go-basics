package main

import "fmt"

func main() {
	var arr [3]int

	arr = fillArray(arr)

	fmt.Println(arr)
}

func fillArray(arr [3]int) [3]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	return arr
}