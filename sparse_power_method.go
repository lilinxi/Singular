package Singular

import (
	"math"
)

func SparsePowerMethod(matrix SparseMatrix, offset, epsilon float64) (eig float64, vec SparseMatrix) {
	if !matrix.IsSquare() {
		panic("")
	}

	if offset != 0 {
		matrix = matrix.Add(SparseMatrixPrototype.Eyes(matrix.Rows()).Scale(offset))
	}

	vec = SparseMatrixPrototype.Full(matrix.Rows(), 1, 1)
	vec.Scale(vec.NormK(2))

	var lastEig float64 = 0
	vec = matrix.Dot(vec)
	lastEig = vec.NormK(2)
	vec = vec.Scale(1 / lastEig)
	for {
		vec = matrix.Dot(vec)
		eig = vec.NormK(2)
		vec = vec.Scale(1 / eig)
		if math.Abs(lastEig-eig) < epsilon {
			break
		}
		lastEig = eig
	}

	if vec.Get(0, 0)*matrix.Dot(vec).Get(0, 0) < 0 {
		eig = -eig
	}

	return eig - offset, vec
}

func SparseInversePowerMethod(matrix SparseMatrix, offset, epsilon float64) (eig float64, vec SparseMatrix) {
	if !matrix.IsSquare() {
		panic("")
	}

	if offset != 0 {
		matrix = matrix.Add(SparseMatrixPrototype.Eyes(matrix.Rows()).Scale(offset))
	}

	vec = SparseMatrixPrototype.Full(matrix.Rows(), 1, 1)
	vec.Scale(vec.NormK(2))

	var lastEig float64 = 0
	vec = LUSolve(matrix, vec)
	lastEig = vec.NormK(2)
	vec = vec.Scale(1 / lastEig)
	for {
		vec = LUSolve(matrix, vec)
		eig = vec.NormK(2)
		vec = vec.Scale(1 / eig)
		if math.Abs(lastEig-eig) < epsilon {
			break
		}
		lastEig = eig
	}

	if vec.Get(0, 0)*matrix.Dot(vec).Get(0, 0) < 0 {
		eig = -eig
	}

	return 1/eig - offset, vec
}
