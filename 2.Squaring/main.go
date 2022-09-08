package main

import (
	"fmt"
	"sync"
)

//2.	Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

var c = make(chan []int)

func main() {
	number := []int{2, 4, 6, 8, 10}

	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(1)

	go squares(number, waitGroup)

	waitGroup.Wait()

	result := <-c
	fmt.Println(result)
}

func squares(numbers []int, waitGroup *sync.WaitGroup) []int {
	sl := []int{}
	var workers = make(chan struct{}, 1) // workers pool (in order to leave the same order)
	wg := new(sync.WaitGroup)            //waitGroup - we need to wait for all goroutines to complete
	for _, number := range numbers {
		workers <- struct{}{}
		wg.Add(1) //one goroutine start to work
		go func(number int) {
			defer func() {
				<-workers
			}()
			defer wg.Done() //one goroutine finished its work
			square := squareCalc(number)
			fmt.Println(square) //stdout
			sl = append(sl, square)
		}(number)
	}
	wg.Wait()
	waitGroup.Done()
	c <- sl
	return sl
}

func consecutive(numbers []int) []int {
	sl := []int{}
	for _, number := range numbers {
		square := squareCalc(number)
		fmt.Println(square)
		sl = append(sl, square)
	}

	return sl
}

func squareCalc(number int) int {
	return number * number
}
