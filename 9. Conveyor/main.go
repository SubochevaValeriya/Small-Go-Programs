package main

import (
	"fmt"
	"sync"
)

// 9.	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	var

}

// read data from slice
func readFromSlice(data []int, ch chan int) {
	defer close(ch)
	for _, elem := range data {
		ch <- elem
	}
}

// read from first slice to second and do
func multi2(chIn chan int, chOut chan int) {
	defer close(chOut)
	for elem := range chIn {
		chOut <- elem * 2
	}
}

//
func readOutput(ch chan int, wg *sync.WaitGroup) {
	for elem := range ch {
		fmt.Println(elem)
		wg.Done() //
	}
}
