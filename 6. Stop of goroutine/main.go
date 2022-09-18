package main

import (
	"context"
	"fmt"
	"time"
)

// 6.	Реализовать все возможные способы остановки выполнения горутины.

func main() {
	closeChannel()
	doneChannel()
	usingContext()
}

func closeChannel() {
	ch := make(chan string, 6)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("Finish")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- "hello"
	ch <- "world"
	close(ch)
	time.Sleep(time.Second)
}

func doneChannel() {
	ch := make(chan string, 6)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch <- "foo":
			case <-done:
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("Received: ", i)
	}

	fmt.Println("Finish")
}

func usingContext() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("foo...")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("Finish")
}
