package main

import "fmt"

// 8.	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	n := 4567
	var i uint = 2
	fmt.Println(setBit(n, i))
	fmt.Println(clearBit(n, i))

}

func setBit(n int, position uint) int {
	n |= 1 << position // shift the number 1 the specified number of spaces in the integer. Then OR it with the original input.
	return n
}

func clearBit(n int, position uint) int {
	mask := ^(1 << position) // shift the number 1 the specified number of spaces in the integer, then flip every bit in the mask with the ^ operator (XOR)
	n &= mask                // use a bitwise AND, which doesn't touch the numbers AND'ed with 1, but which will unset the value in the mask which is set to 0.
	return n
}
