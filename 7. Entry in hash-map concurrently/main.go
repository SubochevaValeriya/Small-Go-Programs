package main

import (
	"fmt"
	"sync"
)

// 7.	Реализовать конкурентную запись данных в map.

func main() {
	var wg sync.WaitGroup
	var m sync.Map

	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func(j int) {
			m.Store(j, fmt.Sprintf("test %v", j))
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Done.")

	for i := 1; i < 10; i++ {
		t, _ := m.Load(i)
		fmt.Println("loading: ", t)
	}
}
