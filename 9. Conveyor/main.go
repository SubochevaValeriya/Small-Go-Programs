package main

import (
	"fmt"
	"sync"
)

// 9.	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

var dataResultForTest []int

func main() {
	generator := func(done <-chan any, integers []int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}

func main() {
	data := []int{1, 4, 3}
	chIn := make(chan int, len(data))
	chOut := make(chan int, len(data))
	wg := sync.WaitGroup{}

	wg.Add(len(data))
	go readFromSlice(data, chIn)
	go multiplicate(chIn, chOut)
	go toOutput(chOut, &wg)

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
		dataResultForTest = append(dataResultForTest, elem)
		fmt.Println(dataResultForTest)
		mu.Unlock()
		wg.Done()
	}
}
