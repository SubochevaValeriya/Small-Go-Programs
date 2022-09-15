package main

import "fmt"

// 12.	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	collection := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(makeSet(collection))

}

func makeSet(collection []string) map[string]struct{} {
	set := map[string]struct{}{} // empty struct because it consumes zero bytes

	for _, str := range collection {
		set[str] = struct{}{}
	}
	return set
}
