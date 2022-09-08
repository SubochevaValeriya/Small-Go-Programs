package main

import (
	"fmt"
	"sync"
)

//3.	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

var c = make(chan int)

func main() {

	number := []int{2, 4, 6, 8, 10}

	c := make(chan int)
	go sumOfSquares(number, c)
	sum := <-c
	fmt.Println(sum)
}

func consecutive(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += squareCalc(number)
	}

	return sum
}

func sumOfSquares(numbers []int, c chan int) int {
	var sum int
	//var workers = make(chan struct{}, 1) //pool of workers
	wg := new(sync.WaitGroup) //waitGroup - we need to wait for all goroutines to complete
	//defer close(workers)
	for _, number := range numbers {
		//	workers <- struct{}{} //
		wg.Add(1) //one goroutine start to work
		go func(number int) {
			defer wg.Done() //one goroutine finished its work
			number = squareCalc(number)
			sum += number //critical section - shared resource
			//			<-workers     //
		}(number)
	}
	wg.Wait()
	c <- sum
	return sum
}

func squareCalc(number int) int {
	return number * number
}
