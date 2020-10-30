package Singular

import (
	"fmt"
	"testing"
)

func TestPowerMethod(t *testing.T) {
	AValue := [][]float64{
		{1, 2, 3},
		{2, 1, 3},
		{3, 3, 5},
	}
	ASparse := NewSparseMatrixFrom2DTable(3, 3, AValue)
	fmt.Println(ASparse)

	eig, vec := PowerMethod(ASparse, Epsilon)
	fmt.Println(eig)
	fmt.Println(vec)
}
