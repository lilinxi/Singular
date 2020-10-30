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
	ADense := NewDenseMatrixFrom2DTable(4, 4, A)
	L, U := LU(ADense)
	fmt.Println("A:", ADense)
	fmt.Println("L:", L)
	fmt.Println("U:", U)
	fmt.Println("L*U:", L.Dot(U))
}
