package main

import (
	"fmt"
	"sync"
)

// 9.	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	data := []int{1, 4, 3}
	var dataResultForTest []int
	chIn := make(chan int, len(data))
	chOut := make(chan int, len(data))
	wg := sync.WaitGroup{}

	wg.Add(len(data))
	go readFromSlice(data, chIn)
	go multiplicate(chIn, chOut)
	go toOutput(chOut, &wg, dataResultForTest)

	fmt.Println(dataResultForTest)

	wg.Wait() // wait until all goroutines will do their job
}

// Read data from slice
func readFromSlice(data []int, ch chan int) {
	defer close(ch)
	for _, elem := range data {
		ch <- elem
	}
}

// Read from first slice to second and do operation x*2
func multiplicate(chIn chan int, chOut chan int) {
	defer close(chOut)
	for elem := range chIn {
		chOut <- elem * 2
	}
}

// Output in stdout
func toOutput(ch chan int, wg *sync.WaitGroup) {
	var mu sync.Mutex
	for elem := range ch {
		fmt.Println(elem)
		mu.Lock()
		dataResultForTest := append(dataResultForTest, elem)
		fmt.Println(dataResultForTest)
		mu.Unlock()
		wg.Done()
	}
}
