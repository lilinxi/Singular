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
