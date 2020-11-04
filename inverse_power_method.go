package Singular

import (
	"math"
)

func InversePowerMethod(matrix SparseMatrix, offset, epsilon float64) (eig float64, vec SparseMatrix) {
	if !matrix.Square() {
		panic("")
	}

	if offset != 0 {
		matrix = matrix.Add(NewSparseMatrixEyes(matrix.Rows()).Scale(offset))
	}

	vec = NewSparseMatrixFull(matrix.Rows(), 1, 1)
	vec.Scale(vec.Norm(2))

	var lastEig float64 = 0
	vec = LUSolve(matrix, vec)
	lastEig = vec.Norm(2)
	vec = vec.Scale(1 / lastEig)
	for {
		vec = LUSolve(matrix, vec)
		eig = vec.Norm(2)
		vec = vec.Scale(1 / eig)
		if math.Abs(lastEig-eig) < epsilon {
			break
		}
		lastEig = eig
	}

	return eig-offset, vec
}
