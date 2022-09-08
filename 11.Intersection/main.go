package main

import "fmt"

//11.	Реализовать пересечение двух неупорядоченных множеств.

func main() {
	m1, m2 := []any{1, 2, 1, 2, 4}, []any{1, 3, 2, 4}

	fmt.Println(intersection(m1, m2))
}

func intersection(m1 []any, m2 []any) any {
	// Create frequency map for first slice
	frequancyMap := map[any]int{}
	result := []any{}

	for _, val := range m1 {
		frequancyMap[val]++
	}

	// check that value from second slice exists in freqMap, append to intersection
	for _, val := range m2 {
		if freq := frequancyMap[val]; freq >= 1 {
			result = append(result, val)
		}
	}

	return result
}
