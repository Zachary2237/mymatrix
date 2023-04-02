package main

import (
	"example/matrix"
	"fmt"
)

func init() {}

func main() {
	m1 := matrix.Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	m2 := matrix.Matrix([][]float64{[]float64{10, 20, 30}, []float64{40, 50, 60}})
	m3, _ := m1.Add(m2)
	fmt.Println(m3)
}
