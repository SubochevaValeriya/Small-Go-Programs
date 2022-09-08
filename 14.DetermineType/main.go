package main

import (
	"fmt"
	"reflect"
)

//14.	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {
	fmt.Println(determineType("golang"))
	fmt.Println(determineType(1911))
	fmt.Println(determineType(true))
	fmt.Println(determineType(make(chan string)))
}

func determineType(val any) string {
	switch val.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case bool:
		return "bool"
	default:
		return reflect.TypeOf(val).String()
	}
	return ""
}
