package main

import (
	"fmt"
)

//23.	Удалить i-ый элемент из слайса.

func main() {

	fmt.Println(deleteFromSlice([]any{"test1", "test2", "test3"}, 2))
}

func deleteFromSlice(sl []any, i int) []any {
	newSl := append(sl[:i], sl[i+1:]...)
	return newSl[:len(sl)-1]
}
