package Singular

import (
	"fmt"
	"testing"
)

func TestLU(t *testing.T) {
	A := [][]float64{
		{2, 1, 1, 0},
		{4, 3, 3, 1},
		{8, 7, 9, 5},
		{6, 7, 9, 8},
	}
	ADense := NewSparseMatrixFrom2DTable(4, 4, A)
	L, U := LU(ADense)
	fmt.Println("A:", ADense)
	fmt.Println("L:", L)
	fmt.Println("U:", U)
	fmt.Println("L*U:", L.Dot(U))

	il := InverseL(L)
	fmt.Println("l:", L)
	fmt.Println("il:", il)
	fmt.Println("il*l:", il.Dot(L))

	// TODO Error
	iu := InverseU(U)
	fmt.Println(iu.Dot(U))
}

func TestLUSolve(t *testing.T) {
	A := NewSparseMatrixFrom2DTable(3, 3, [][]float64{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 11},
	})
	b := NewSparseMatrixFrom1DList([]float64{1, 1, 1})
	L, U := LU(A)
	fmt.Println("L:", L)
	fmt.Println("U:", U)

	x := LUSolve(A, b)
	fmt.Println(x)
}
