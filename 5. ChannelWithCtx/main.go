package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

// 5.	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	channel(ctx)

}

func channel(ctx context.Context) {
	ch := make(chan int, 1)

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			os.Exit(200)
		default:
			sent := i
			ch <- sent
			received := <-ch
			fmt.Println(received)
		}

	}
}
