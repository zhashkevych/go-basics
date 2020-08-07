package main

import "fmt"

func main() {
	todoList := [4]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
		"дописать 3-ю часть серии уроков",
	}

	tasks := todoList[1:4] 
	for i := range tasks {
		fmt.Println(tasks[i])
	}

	fmt.Println("\n---- ПОСЛЕ changеTasks() ----\n")
	changеTasks(tasks)

	for i := range tasks {
		fmt.Println(tasks[i])
	}

	fmt.Println("\n---- МАССИВ todoList ПОСЛЕ changеTasks() ----\n")
	fmt.Println(todoList)
}

func changеTasks(tasks []string) {
	tasks[0] = "пройти курс по Go"
	tasks[1] = "сказать автору спасибо :)"
}