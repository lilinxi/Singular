package Singular

import (
	"fmt"
	"testing"
)

func TestSparseMatrix(t *testing.T) {
	a := NewSparseMatrixFrom1DList([]float64{1, 2, 3, 4})
	b := NewSparseMatrixFrom1DList([]float64{1, 2, 3, 4}).Transpose()
	ab := a.Dot(b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(ab)

	abR0 := ab.GetRow(0)
	abR1 := ab.GetRow(1)
	abR2 := ab.GetRow(2)
	abR3 := ab.GetRow(3)
	fmt.Println(abR0)
	fmt.Println(abR1)
	fmt.Println(abR2)
	fmt.Println(abR3)

	abC0 := ab.GetCol(0)
	abC1 := ab.GetCol(1)
	abC2 := ab.GetCol(2)
	abC3 := ab.GetCol(3)
	fmt.Println(abC0)
	fmt.Println(abC1)
	fmt.Println(abC2)
	fmt.Println(abC3)

	fmt.Println(ab)
	ab.SetRow(0, abR3)
	ab.SetCol(0, abC3)
	fmt.Println(ab)

	fmt.Println(NewSparseMatrixEyes(4))
	fmt.Println(ab.Add(NewSparseMatrixEyes(4).Scale(-4)))
	fmt.Println(ab.NormInf())

}
