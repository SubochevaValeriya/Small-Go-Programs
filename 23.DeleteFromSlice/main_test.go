package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	sl       []any
	i        int
	expected []any
}

func TestDeleteFromSlice(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with strings", sl: []any{"test1", "test2", "test3"}, i: 2, expected: []any{"test1", "test2"}},
		{Name: "case with ints", sl: []any{0, 1, 2, 3, 4, 5}, i: 3, expected: []any{0, 1, 2, 4, 5}},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := deleteFromSlice(cse.sl, cse.i)
			assert.Equalf(t, cse.expected, result, "for %s and %v expected %s, got %s", cse.sl, cse.i, cse.expected, result)
		})
	}
}
