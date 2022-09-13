package main

import (
	"fmt"
)

//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(wg sync.WaitGroup, i int) {
//			fmt.Println(i)
//			wg.Done()
//		}(wg, i)
//	}
//	wg.Wait()
//	fmt.Println("exit")
//}

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
