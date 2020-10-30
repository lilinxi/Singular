package Singular

import (
	"fmt"
	"testing"
)

func TestPowerMethod(t *testing.T) {
	AValue := [][]complex128{
		{1, 2, 3},
		{2, 1, 3},
		{3, 3, 5},
	}
	ASparse := NewSparseMatrixFrom2DTable(3, 3, AValue)
	fmt.Println(ASparse)

	vec, eig := PowerMethod(ASparse, Epsilon)
	fmt.Println(vec)
	fmt.Println(eig)
}
