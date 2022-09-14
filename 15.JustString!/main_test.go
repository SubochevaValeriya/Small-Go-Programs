package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name      string
	a         int
	b         int
	expectedA int
	expectedB int
}

func TestChangePlacesGolang(t *testing.T) {
	testCases := []TestCase{
		{Name: "case 3, 4", a: 3, b: 4, expectedA: 4, expectedB: 3},
		{Name: "case 1000000, 1", a: 1000000, b: 1, expectedA: 1, expectedB: 1000000},
		{Name: "case -10, 0", a: -10, b: 0, expectedA: 0, expectedB: -10},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			a, b := changePlacesGolang(cse.a, cse.b)
			assert.Equalf(t, combineValues(cse.expectedA, cse.expectedB), combineValues(a, b), "for %d and %d expected %t, got %t", cse.a, cse.b, cse.expectedA, a)
		})
	}
}

func combineValues(a, b any) []any {
	return []any{a, b}
}

func TestChangePlacesMethod2(t *testing.T) {
	testCases := []TestCase{
		{Name: "case 3, 4", a: 3, b: 4, expectedA: 4, expectedB: 3},
		{Name: "case 1000000, 1", a: 1000000, b: 1, expectedA: 1, expectedB: 1000000},
		{Name: "case -10, 0", a: -10, b: 0, expectedA: 0, expectedB: -10},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			a, b := changePlacesMethod2(cse.a, cse.b)
			assert.Equalf(t, combineValues(cse.expectedA, cse.expectedB), combineValues(a, b), "for %d and %d expected %t, got %t", cse.a, cse.b, cse.expectedA, a)
		})
	}
}
