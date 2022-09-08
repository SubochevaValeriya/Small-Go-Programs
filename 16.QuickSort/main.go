package main

import "fmt"

//16.	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

var justString string

func main() {
	sl := []int{11, 1, -2, 213, -20, 9}

	fmt.Println(quickSort(sl))
}

// time complexity: O(n*log n) - O(n) - division an array into two parts, O(log n) - recursion depth (O(n) in bad case)
func quickSort(sl []int) []int {
	// Choose of base element (median of third)
	base, baseIndx := chooseBaseElement(sl)
	left, right := 0, len(sl)-1

	// Move base element to the right
	sl[baseIndx], sl[right] = sl[right], sl[baseIndx]

	// Go through the slice and put smaller elements to the left part
	for i := 0; i < len(sl); i++ {
		if sl[i] < base {
			sl[i], sl[left] = sl[left], sl[i]
			left++
		}
	}

	// Put base element between smaller and higher elements (after last smaller element)
	sl[left], sl[right] = sl[right], sl[left]

	// Recursively use quick sort algorithm for left and right parts (if their length > 1)
	if len(sl[:left]) > 1 {
		quickSort(sl[:left])
	}

	if len(sl[left+1:]) > 1 {
		quickSort(sl[left+1:])
	}

	return sl
}

func chooseBaseElement(sl []int) (int, int) {
	// if len == 2, just choose first element
	if len(sl) < 3 {
		return sl[0], 0
	}

	//choosing median from first, mid and end elements
	first, mid, end := sl[0], sl[len(sl)/2], sl[len(sl)-1]

	if (first > end && first < mid) || (first > mid && first < end) {
		return first, 0
	}

	if (mid > end && first < mid) || (mid > first && mid < end) {
		return mid, len(sl) / 2
	}

	if (end > first && end < mid) || (end > mid && end < first) {
		return end, len(sl) - 1
	}
	return sl[0], 0
}
