package main

import "fmt"

// Определение структуры Human
type Human struct {
	Name string
	Age  int
}

// Метод для структуры Human
func (h Human) printName() {
	fmt.Printf("Hello! My name is %s.\n", h.Name)
}

// Определение структуры Action, в которой встроена структура Human
type Action struct {
	Human
}

// Метод для структуры Action
func (a Action) DoSomething() {
	fmt.Printf("%s is doing something.\n", a.Name)
}

func main() {
	// Создание объекта структуры Action
	person := Action{
		Human: Human{Name: "Maxim", Age: 22},
	}

	person.printName()   // Вызов метода, унаследованного от структуры Human
	person.DoSomething() // Вызов метода, определенного в структуре Action
}
