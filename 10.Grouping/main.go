package main

import "fmt"

// 10.	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(grouping(data))
}

func grouping(data []float64) map[int][]float64 {
	m := map[int][]float64{}
	for _, val := range data {
		nearTen := int(val) - int(val)%10
		m[nearTen] = append(m[nearTen], val)
	}

	return m
}
