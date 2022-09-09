package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// 7.	Реализовать конкурентную запись данных в map.

var m sync.Map
var wait = make(chan struct{}, 1)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		sigChan := <-sigChan
		fmt.Println(sigChan)
		fmt.Println("ff")
		ordinaryMap := map[string]any{}
		m.Range(func(key, value any) bool {
			ordinaryMap[fmt.Sprint(key)] = value
			return false
		})
		<-wait
		os.Exit(1)
		// поймали один из
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString(byte('\n'))
		keyAndValue := strings.Split(text, ",")
		//if len(keyAndValue) != 2 {
		//	fmt.Println("incorrect input format")
		//	return
		//}
		go concurrentMap(keyAndValue[0], keyAndValue[1])
	}

	//var wg sync.WaitGroup
	//var m sync.Map
	//
	//wg.Add(10)
	//for i := 1; i <= 10; i++ {
	//	go func(j int) {
	//		m.Store(j, fmt.Sprintf("test %v", j))
	//		wg.Done()
	//	}(i)
	//}
	//
	//wg.Wait()
	//
	////fmt.Println(m)
	//fmt.Println("Done.")
	//
	//for i := 1; i < 10; i++ {
	//	t, _ := m.Load(i)
	//	fmt.Println("loading: ", t)
	//}
}

func concurrentMap(key, value any) {
	wait <- struct{}{}

	var wg sync.WaitGroup

	wg.Add(1)
	//go func(j int) {
	m.Store(key, value)
	fmt.Printf("stored: [%v:%v]", key, value)
	wg.Done()

	wg.Wait()

	//fmt.Println(m)
	fmt.Println("Done.")

}
