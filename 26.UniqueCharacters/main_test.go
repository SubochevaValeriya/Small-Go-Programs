package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	str      string
	expected bool
}

func TestReverseString(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with abcd", str: "abcd", expected: true},
		{Name: "case with abCdefAaf", str: "abCdefAaf", expected: false},
		{Name: "case with привет", str: "привет", expected: true},
		{Name: "case with пЕПел", str: "пЕПел", expected: false},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := unique(cse.str)
			assert.Equalf(t, cse.expected, result, "for %s expected %s, got %s", cse.str, cse.expected, result)
		})
	}
}
