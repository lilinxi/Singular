package Singular

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestGaussElimination(t *testing.T) {
	AValue := [][]float64{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 11},
	}
	ASparse := SparseMatrixPrototype.From2DTable(AValue)

	bValue := [][]float64{{1}, {1}, {1}}
	bSparse := SparseMatrixPrototype.From2DTable(bValue)

	x := SparseGaussSolve(ASparse, bSparse)
	assert.T(t, EqualSparseMatrix(x, SparseMatrixPrototype.From1DList(
		[]float64{
			-0.33333333333333326,
			0.3333333333333333,
			0},
	)))
}
