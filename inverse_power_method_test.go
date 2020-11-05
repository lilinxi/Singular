package Singular

import (
	"fmt"
	"testing"
)

func TestInversePowerMethod(t *testing.T) {
	A := NewSparseMatrixFrom2DTable(3, 3, [][]float64{
		{1, 2, 3},
		{2, 1, 3},
		{3, 3, 5},
	})
	fmt.Println(A)
	//ans =
	//
	//-1.00000
	//-0.35890
	//8.35890

	eig, vec := InversePowerMethod(A, 0, Epsilon)
	fmt.Println(eig)
	fmt.Println(vec)

	fmt.Println(vec.Scale(eig))
	fmt.Println(A.Dot(vec))

	for i := -10.0; i < 10; i+=1.05 {
		eig, _ := InversePowerMethod(A, i, Epsilon)
		fmt.Println(eig)
	}
}
