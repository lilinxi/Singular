package Singular

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func TestLU(t *testing.T) {
	A := [][]float64{
		{2, 1, 1, 0},
		{4, 3, 3, 1},
		{8, 7, 9, 5},
		{6, 7, 9, 8},
	}
	ASpare := SparseMatrixPrototype.From2DTable(A)
	L, U := LU(ASpare)
	assert.T(t, EqualSparseMatrix(L, SparseMatrixPrototype.From2DTable(
		[][]float64{
			{1, 0, 0, 0},
			{2, 1, 0, 0},
			{4, 3, 1, 0},
			{3, 4, 1, 1},
		},
	)))
	assert.T(t, EqualSparseMatrix(U, SparseMatrixPrototype.From2DTable(
		[][]float64{
			{2, 1, 1, 0},
			{0, 1, 1, 1},
			{0, 0, 2, 2},
			{0, 0, 0, 2},
		},
	)))
	assert.T(t, EqualSparseMatrix(L.Dot(U), ASpare))

	il := InverseL(L)
	assert.T(t, EqualSparseMatrix(il, SparseMatrixPrototype.From2DTable(
		[][]float64{
			{1, 0, 0, 0},
			{-2, 1, 0, 0},
			{2, -3, 1, 0},
			{3, -1, -1, 1},
		},
	)))
	assert.T(t, EqualSparseMatrix(L.Dot(il), SparseMatrixPrototype.Eyes(4)))
}

func TestLUSolve(t *testing.T) {
	A := SparseMatrixPrototype.From2DTable([][]float64{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 11},
	})
	b := SparseMatrixPrototype.From1DList([]float64{1, 1, 1})
	L, U := LU(A)
	fmt.Println("L:", L)
	fmt.Println("U:", U)

	x := LUSolve(A, b)
	assert.T(t, EqualSparseMatrix(x, SparseMatrixPrototype.From1DList(
		[]float64{
			-0.33333333333333326,
			0.3333333333333333,
			0},
	)))
}
