package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	A        Point
	B        Point
	Expected float64
}

func TestDistance(t *testing.T) {
	testCases := []TestCase{
		{Name: "case with 2.28", A: Point{x: 1.2, y: 3.4}, B: Point{3.0, 2.0}, Expected: 2.280350850198276},
		{Name: "case with 0", A: Point{x: 0.0, y: 0.0}, B: Point{0.0, 0.0}, Expected: 0.00},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result := Distance(cse.A, cse.B)
			assert.Equalf(t, cse.Expected, result, "for %f and %f expected %f, got %f", cse.A, cse.B, cse.Expected, result)
		})
	}
}
