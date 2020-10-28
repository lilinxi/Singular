package Singular

import (
	"fmt"
	"testing"
)

func TestSparseMatrix(t *testing.T) {
	values := make(map[int]map[int]complex128)
	values[0] = make(map[int]complex128)
	values[0][0] = 1
	values[0][1] = 2
	values[1] = make(map[int]complex128)
	values[1][0] = 3
	values[1][1] = 4

	matrix := NewSparseMatrix(2, 2, values)
	fmt.Println(matrix)

	fmt.Println(matrix.Add(matrix))
	fmt.Println(matrix.Sub(matrix))
	fmt.Println(matrix.Dot(matrix))
	fmt.Println(matrix.Scale(2))
	fmt.Println(matrix.Transpose())

	values = make(map[int]map[int]complex128)
	values[0] = make(map[int]complex128)
	values[0][0] = 1
	values[1] = make(map[int]complex128)
	values[1][0] = 4
	values[2] = make(map[int]complex128)
	values[2][1] = 5

	matrix = NewSparseMatrix(3, 3, values)
	fmt.Println(matrix)

	fmt.Println(matrix.Add(matrix))
	fmt.Println(matrix.Sub(matrix))
	fmt.Println(matrix.Dot(matrix))
	fmt.Println(matrix.Scale(2))
	fmt.Println(matrix.Transpose())
}
