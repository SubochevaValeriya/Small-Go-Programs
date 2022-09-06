package main

import "fmt"

//13.	Поменять местами два числа без создания временной переменной.

func main() {

	fmt.Println(changePlacesGolang(3, 4))
	fmt.Println(changePlacesMethod2(3, 4))
}

func changePlacesGolang(a, b int) (int, int) {
	a, b = b, a

	return a, b
}

func changePlacesMethod2(a, b int) (int, int) {
	a = b - a
	b = b - a
	a = b + a

	return a, b

}
