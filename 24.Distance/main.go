package main

import (
	"fmt"
	"math"
)

//23.	Удалить i-ый элемент из слайса.

type Point struct {
	x float64
	y float64
}

func NewPointConstructor(x, y float64) Point {
	return Point{x: x, y: y}
}

func Distance(pA Point, pB Point) float64 {
	AC := pB.x - pA.x
	BC := pB.y - pA.y
	AB := math.Sqrt(AC*AC + BC*BC)
	return AB
}

func main() {

	fmt.Printf("%.2f", Distance(NewPointConstructor(1.2, 3.4), NewPointConstructor(3.0, 2.0)))
}
