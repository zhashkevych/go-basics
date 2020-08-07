package main

import (
	"fmt"
	"github.com/zhashkevych/go-basics/employees/storage"

)

func main() {
	ms := storage.NewMemoryStorage()
	ds :=  storage.NewDumbStorage()

	fmt.Println(ms.data)

	spawnEmployees(ms)
	fmt.Println(ms.Get(3))

	spawnEmployees(ds)
}

func spawnEmployees(s storage.Storage) {
	for i := 1; i <= 10; i++ {
		s.Insert(storage.Employee{Id: i})
	}
}