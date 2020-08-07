package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // время старта

	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	result1 := make(chan int)
	result2 := make(chan int)

	// очень сложные вычисления
	go calculateSomething(1000, result1)

	// а эти вычисления еще сложнее
	go calculateSomething(2000, result2)

	fmt.Println(<-result1)
	fmt.Println(<-result2)

	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))
}

func calculateSomething(n int, res chan int) {
	t := time.Now()

	result := 0
	for i := 0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Время выполнения расчетов: %s\n", time.Since(t))
	res <- result
}
