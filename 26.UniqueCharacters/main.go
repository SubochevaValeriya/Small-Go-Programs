package main

import (
	"fmt"
	"strings"
)

//26.	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func main() {

	str := "abcdB"

	fmt.Println(unique(str))
}

func unique(str string) bool {
	str = strings.ToLower(str) // case insensitive

	// Create map and check in the loop that our symbol occurs only once (use rune in order to not be limited to latin only)
	m := map[rune]int{}

	for _, symbol := range []rune(str) {
		if m[symbol] > 0 {
			return false
		}
		m[symbol]++
	}
	return true
}
