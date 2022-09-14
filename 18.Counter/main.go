package main

import (
	"fmt"
	"sync"
)

//18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	counter int
	mu      sync.Mutex
	wg      sync.WaitGroup
}

func main() {

	counter := Counter{
		counter: 0,
		mu:      sync.Mutex{},
		wg:      sync.WaitGroup{},
	}

	counter.count(100) // count to n
}

func (c *Counter) count(n int) {
	c.wg.Add(n)
	for i := 0; i < n; i++ {
		c.mu.Lock()
		c.counter++
		fmt.Println("I'm counting:", c.counter)
		c.mu.Unlock()
	}
	fmt.Println("Result:", c.counter)
}
