package main

import (
	"fmt"
	"strings"
)

//21.	Реализовать паттерн «адаптер» на любом примере

func main() {

	fmt.Println(reverseString("snow dog sun"))
}

func reverseString(str string) string {
	var reversedString string
	strSlice := strings.Split(str, " ")
	for i := len(strSlice) - 1; i >= 0; i-- {
		reversedString += strSlice[i]
		if i != 0 {
			reversedString += " "
		}
	}
	return reversedString
}
