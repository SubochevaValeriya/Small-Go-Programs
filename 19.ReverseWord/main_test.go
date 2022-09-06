package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	word     string
	expected string
}

func TestReverseWord(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with главрыба", word: "главрыба", expected: "абырвалг"},
		{Name: "case with hello 😺", word: "hello 😺", expected: "😺 olleh"},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := reverseWord(cse.word)
			assert.Equalf(t, cse.expected, result, "for %s expected %s, got %s", cse.word, cse.expected, result)
		})
	}
}
