package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

type TestCase struct {
	Name     string
	slice    []int
	expected []int
}

func TestSquares(t *testing.T) {
	testCases := []TestCase{
		{Name: "case [2, 4, 6, 8, 10]", slice: []int{2, 4, 6, 8, 10}, expected: consecutive([]int{2, 4, 6, 8, 10})},
		{Name: "case [-1, 10000, 20, -333, 99, 0, 455, 764, 3]", slice: []int{-1, 10000, 20, -333, 99, 0, 455, 764, 3}, expected: consecutive([]int{-1, 10000, 20, -333, 99, 0, 455, 764, 3})},
		{Name: "case [1, 1, 1, 2, 1, 0, 2, 1, 3, 1, 2, 2, 1, 2, 0, 1, 0, 1, 2, 1, 2, 1, 1]", slice: []int{1, 1, 1, 2, 1, 0, 2, 1, 3, 1, 2, 2, 1, 2, 0, 1, 0, 1, 2, 1, 2, 1, 1}, expected: consecutive([]int{1, 1, 1, 2, 1, 0, 2, 1, 3, 1, 2, 2, 1, 2, 0, 1, 0, 1, 2, 1, 2, 1, 1})},
	}
	for _, cse := range testCases {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			waitGroup := new(sync.WaitGroup)
			waitGroup.Add(1)
			go squares(cse.slice, waitGroup)
			waitGroup.Wait()
			result := <-c
			assert.Equalf(t, cse.expected, result, "for %d expected %t, got %t", cse.slice, cse.expected)
		})
	}
}
