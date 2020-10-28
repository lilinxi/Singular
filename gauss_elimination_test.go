package Singular

import (
	"fmt"
	"testing"
)

func TestGaussElimination(t *testing.T) {
	AValue := [][]complex64{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 11},
	}
	ASparse := NewSparseMatrixFrom2DTable(3, 3, AValue)
	fmt.Println(ASparse)

	bValue := [][]complex64{{1}, {1}, {1}}
	bSparse := NewSparseMatrixFrom2DTable(3, 1, bValue)
	fmt.Println(bSparse)

	x := GaussElimination(ASparse, bSparse)
	fmt.Println(x) // {-1/3, 1/3, 0}
}
