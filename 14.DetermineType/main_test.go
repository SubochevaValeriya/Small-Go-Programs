package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	val      any
	expected string
}

func TestDetermineType(t *testing.T) {
	testCases := []TestCase{
		{Name: "case string", val: "cat", expected: "string"},
		{Name: "case int", val: 1, expected: "int"},
		{Name: "case bool", val: true, expected: "bool"},
		{Name: "case channel", val: make(chan string), expected: "chan string"},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := determineType(cse.val)
			assert.Equalf(t, cse.expected, result, "for %v expected %v, got %v", cse.val, cse.expected, result)
		})
	}
}
