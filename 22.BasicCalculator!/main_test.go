package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	str      string
	expected string
}

func TestReverseString(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with snow dog sun", str: "snow dog sun", expected: "sun dog snow"},
		{Name: "case with Test Not If", str: "Test Not If", expected: "If Not Test"},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := reverseString(cse.str)
			assert.Equalf(t, cse.expected, result, "for %s expected %s, got %s", cse.str, cse.expected, result)
		})
	}
}
