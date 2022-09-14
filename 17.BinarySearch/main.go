package main

import (
	"errors"
	"fmt"
)

//17.	Реализовать бинарный поиск встроенными методами языка.

func main() {
	data := []int{1, 2, 3, 11, 35, 55}
	searchElement := 11

	fmt.Println(binarySerch(data, searchElement))
}

// time complexity: O(log n)
func binarySerch(data []int, searchElement int) (int, error) {

	// Left, right elements

	left, right := 0, len(data)-1

	for left <= right {
		// Finding the position of the middle element
		m := (right + left) / 2
		// If we find target value
		if data[m] == searchElement {
			return m, nil
		}

		// Changing interval
		if data[m] < searchElement {
			left = m + 1
		} else {
			right = m - 1
		}
	}

	// If not found, return error

	return 0, errors.New("not found")
}
