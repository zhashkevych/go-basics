package main

import "fmt"

type employee struct{
	name string
	sex string // пол
	age int 
	salary int // зарплата
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name: name,
		sex: sex,
		age: age,
		salary: salary,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник: %s\nВозраст: %d\nЗарплата: %d\n", e.name, e.age, e.salary)
}

func (e employee) setName(name string) {
	e.name = name
}

func main() {
	employee1 := newEmployee("Вася", "М", 25, 1500)
	employee2 := newEmployee("Петя", "М", 28, 2000)

	employee1.setName("Генадий")

	fmt.Printf("%s\n", employee1.getInfo())
	fmt.Printf("%s", employee2.getInfo())
}

