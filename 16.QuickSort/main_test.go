package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	sl       []int
	expected []int
}

func TestGrouping(t *testing.T) {
	testCases := []TestCase{
		{Name: "case [11, 1, -2, 213, -20, 9]", sl: []int{11, 1, -2, 213, -20, 9}, expected: []int{-20, -2, 1, 9, 11, 213}},
		{Name: "case [747747, -1, 2, -1, 4, 0, -1, 2, 2,4, 1, -5]", sl: []int{747747, -1, 2, -1, 4, 0, -1, 2, 2, 4, 1, -5}, expected: []int{-5, -1, -1, -1, 0, 1, 2, 2, 2, 4, 4, 747747}},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := quickSort(cse.sl)
			assert.Equalf(t, cse.expected, result, "for %d expected %t, got %t", cse.sl, cse.expected, result)
		})
	}
}
