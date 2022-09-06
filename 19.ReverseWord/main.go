package main

import (
	"fmt"
)

//19.	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {

	fmt.Println(reverseWord("Test Тест 😺"))
}

func reverseWord(word string) string {
	var reversedWord string
	wordRune := []rune(word) // take into account that symbols can take more than one byte

	for i := len(wordRune) - 1; i >= 0; i-- {
		reversedWord += string(wordRune[i])
	}

	return reversedWord
}
