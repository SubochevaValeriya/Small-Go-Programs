package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	m1       []any
	m2       []any
	expected []any
}

func TestReverseString(t *testing.T) {
	testCases := []TestCase{
		{Name: "case 1", m1: []any{1, 2, 1, 2, 4, 3}, m2: []any{1, 5, 4, 2, 7}, expected: []any{1, 4, 2}},
		{Name: "case 2", m1: []any{"h", "rr", "or"}, m2: []any{"g", "p", "rr"}, expected: []any{"rr"}},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := intersection(cse.m1, cse.m2)
			assert.Equalf(t, cse.expected, result, "for %s & %s expected %s, got %s", cse.m1, cse.m2, cse.expected, result)
		})
	}
}
