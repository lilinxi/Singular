package Singular

import "math"

const Epsilon = 1e-12

func Equal(a float64, b float64) bool {
	return math.Abs(a-b) < Epsilon
}

func EqualSparseMatrix(a SparseMatrix, b SparseMatrix) bool {
	if a.Rows() != b.Rows() || a.Cols() != b.Cols() {
		panic("")
	}
	for i := 0; i < a.Rows(); i++ {
		for j := 0; j < a.Cols(); j++ {
			if !Equal(a.Get(i, j), b.Get(i, j)) {
				return false
			}
		}
	}
	return true
}
