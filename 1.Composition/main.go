package main

import "fmt"

//1.	Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	Name string
	Age  int
	Action
}

func (h Human) sayHello() {
	fmt.Println("Hello")
}

type Action struct {
	Act string
}

func (a Action) act() {
	fmt.Println("acting")
}

func main() {
	action := Action{Act: "act"}
	person := Human{
		Name:   "",
		Age:    0,
		Action: action,
	}

	person.sayHello()
	person.act()
}
