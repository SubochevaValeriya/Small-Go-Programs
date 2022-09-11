package main

import "fmt"

// 4.	Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала
//и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	data := []any{1, 2, 3}

	ch := make(chan any, len(data))

	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}

	fmt.Println(workers(len(data), ch))
}

func workers(n int, ch <-chan any) []any {
	data := []any{}
	workers := make(chan struct{}, 1)
	for i := 1; i <= n; i++ {
		fmt.Println("h")
		go func() {
			fmt.Println("h")
			fmt.Println(ch)
			data = append(data, ch)
			workers <- struct{}{}
		}()
	}
	return data
}
