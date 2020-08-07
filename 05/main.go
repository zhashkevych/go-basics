package main

import (
	"errors"
	"fmt"
)

type employee struct {
	id     int
	name   string
	age    int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e

	return nil
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("employee with such id doesn't exist")
	}

	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

type dumbStorage struct{}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) insert(e employee) error {
	fmt.Printf("вставка пользователя с id: %d прошла успешно\n", e.id)
	return nil
}

func (s *dumbStorage) get(id int) (employee, error) {
	e := employee{
		id: id,
	}

	return e, nil
}

func (s *dumbStorage) delete(id int) error {
	fmt.Printf("удаление пользователя с id: %d прошло успешно\n", id)
	return nil
}

func main() {
	ms := newMemoryStorage()
	ds := newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))

	spawnEmployees(ds)
}

func spawnEmployees(s storage) {
	for i := 1; i <= 10; i++ {
		s.insert(employee{id: i})
	}
}
