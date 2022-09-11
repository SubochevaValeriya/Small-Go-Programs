package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	data     []int
	expected []int
}

func TestGrouping(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with 1, 4, 3", data: []int{1, 4, 3}, expected: []int{2, 8, 6}},
		{Name: "case with 1, 4, 3", data: []int{1, 2, 1000, 3, 4, 12, 1}, expected: []int{2, 4, 2000, 6, 8, 24, 2}},
	}
	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := pipeline(cse.data)
			assert.Equalf(t, cse.expected, result, "for %d expected %t, got %t", cse.data, cse.expected, result)
		})
	}
}
