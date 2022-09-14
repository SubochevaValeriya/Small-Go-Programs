package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name          string
	data          []int
	searchElement int
	expected      int
	expectedError error
}

func TestGrouping(t *testing.T) {
	testCases := []TestCase{
		{Name: "case [1, 2, 3, 11, 35, 55]", data: []int{1, 2, 3, 11, 35, 55}, searchElement: 11, expected: 3, expectedError: nil},
		{Name: "case [1, 2, 3, 11, 35, 55]", data: []int{1, 2, 3, 11, 35, 55}, searchElement: 4, expected: 0, expectedError: errors.New("not found")},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result, err := binarySerch(cse.data, cse.searchElement)
			assert.Equalf(t, cse.expected, result, "for %d & %d expected %t, got %t", cse.data, cse.searchElement, cse.expected, result)
			if cse.expectedError != nil {
				assert.Error(t, cse.expectedError, err, "for %d & %d expected %t, got %t", cse.data, cse.searchElement, cse.expected, result)
			}
		})
	}
}
