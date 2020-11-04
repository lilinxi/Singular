package Singular

import (
	"fmt"
	"testing"
)

func TestSparseMatrix(t *testing.T) {
	values := make(map[int]map[int]float64)
	values[0] = make(map[int]float64)
	values[0][0] = 1
	values[0][1] = 2
	values[1] = make(map[int]float64)
	values[1][0] = 3
	values[1][1] = 4

	matrix := NewSparseMatrixFromMap(2, 2, values)
	fmt.Println(matrix)

	fmt.Println(matrix.Add(matrix))
	fmt.Println(matrix.Sub(matrix))
	fmt.Println(matrix.Dot(matrix))
	fmt.Println(matrix.Scale(2))
	fmt.Println(matrix.Transpose())

	values = make(map[int]map[int]float64)
	values[0] = make(map[int]float64)
	values[0][0] = 1
	values[1] = make(map[int]float64)
	values[1][0] = 4
	values[2] = make(map[int]float64)
	values[2][1] = 5

	matrix = NewSparseMatrixFromMap(3, 3, values)
	fmt.Println(matrix)

	fmt.Println(matrix.Add(matrix))
	fmt.Println(matrix.Sub(matrix))
	fmt.Println(matrix.Dot(matrix))
	fmt.Println(matrix.Scale(2))
	fmt.Println(matrix.Transpose())

	values = make(map[int]map[int]float64)
	values[0] = make(map[int]float64)
	values[0][0] = 1
	values[0][1] = 2
	values[0][2] = 3
	matrix = NewSparseMatrixFromMap(3, 3, values)
	fmt.Println(matrix)
	fmt.Println(matrix.Norm2Square())
}
