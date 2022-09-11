package main

import (
	"fmt"
)

// 9.	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

var dataResultForTest []int

func main() {
	integers := []int{1, 4, 3}
	fmt.Println(pipeline(integers))
}

func pipeline(integers []int) []int {
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

	done := make(chan any)
	defer close(done)

	intStream := generator(done, integers)
	pipelineChan := multiply(done, intStream, 2)

	result := []int{}
	for v := range pipelineChan {
		result = append(result, v)
		fmt.Println(v)
	}
	return result
}
