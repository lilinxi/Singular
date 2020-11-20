package Singular

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

/**
构造 Householder 矩阵可以将任何向量投影到 (-σ,0,0,0……)^T
*/
func TestDenseHouseholder(t *testing.T) {
	x := DenseMatrixPrototype.From1DList([]float64{3, 4, 12})
	y := DenseMatrixPrototype.From1DList([]float64{-13, 0, 0})
	w := x.Sub(y).Normal()
	H := DenseHouseholder(w)
	fmt.Println(H)
	fmt.Println(H.Scale(13))
	fmt.Println(H.Dot(x))
	fmt.Println(EqualDenseMatrix(H.Dot(x), y))
}

func TestDenseQR(t *testing.T) {
	A := DenseMatrixPrototype.From2DTable(
		[][]float64{
			{2, 1, 0},
			{1, 1, 1},
			{2, 0, 2},
		})
	Q, R := DenseQR(A)

	fmt.Println(Q)
	fmt.Println(R)

	assert.T(t, EqualDenseMatrix(Q, DenseMatrixPrototype.From2DTable(
		[][]float64{
			{-2.0 / 3.0, -1.0 / 3.0, -2.0 / 3.0},
			{-1.0 / 3.0, -2.0 / 3.0, 2.0 / 3.0},
			{-2.0 / 3.0, 2.0 / 3.0, 1.0 / 3.0},
		},
	)))
	assert.T(t, EqualDenseMatrix(R, DenseMatrixPrototype.From2DTable(
		[][]float64{
			{-3.0, -1.0, -5.0 / 3.0},
			{0.0, -1.0, 2.0 / 3.0},
			{0.0, 0.0, 4.0 / 3.0},
		},
	)))
}
