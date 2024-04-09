package main

import "fmt"

// Определение структуры Human
type Human struct {
	Name   string
	Age    int
	Gender string
}

// Метод Speak для структуры Human
func (h *Human) Speak() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

// Определение структуры Action, встраивающей Human
type Action struct {
	Human // встраивание структуры Human
	Hobby string
}

// Метод DoAction для структуры Action
func (a *Action) DoAction() {
	fmt.Printf("%s is %s and enjoys %s.\n", a.Name, a.Gender, a.Hobby)
}

func main() {
	// Создание объекта Action
	action := Action{
		Human: Human{
			Name:   "Alice",
			Age:    30,
			Gender: "female",
		},
		Hobby: "reading",
	}

	// Вызов методов встроенной структуры Human и Action
	action.Speak()
	action.DoAction()
}
