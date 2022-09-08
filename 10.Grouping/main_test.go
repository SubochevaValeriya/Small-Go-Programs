package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	data     []float64
	expected map[int][]float64
}

func TestGrouping(t *testing.T) {
	testCases := []TestCase{
		{Name: "case []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}", data: []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}, expected: map[int][]float64{-20: {-25.4, -27, -21}, 10: {13, 19, 15.5}, 20: {24.5}, 30: {32.5}}},
		{Name: "case []float64{-1.0, 100.0, -1.1, -1.9, 12.5, 12.7, 12848328428}", data: []float64{-1.0, 100.0, -1.1, -1.9, 12.5, 12.7, 12848328428.3, -13828302.1}, expected: map[int][]float64{0: {-1.0, -1.1, -1.9}, 100: {100.0}, 10: {12.5, 12.7}, 12848328420: {12848328428.3}, -13828300: {-13828302.1}}},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := grouping(cse.data)
			assert.Equalf(t, cse.expected, result, "for %d expected %t, got %t", cse.data, cse.expected, result)
		})
	}
}
