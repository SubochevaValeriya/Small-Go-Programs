package main

import (
	"fmt"
	"strings"
)

//20.	Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {

	fmt.Println(reverseString("snow dog sun"))
}

func reverseString(str string) string {
	var reversedString string
	strSlice := strings.Split(str, " ")
	for i := len(strSlice) - 1; i >= 0; i-- {
		if i != 0 {
			reversedString += strSlice[i] + " "
		} else {
			reversedString += strSlice[i]
		}
	}
	return reversedString
}
