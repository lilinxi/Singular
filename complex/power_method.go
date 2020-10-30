package Singular

import (
	"math/cmplx"
)

func PowerMethod(matrix SparseMatrix, epsilon float64) (vec SparseMatrix, eig complex128) {
	values := make(map[int]map[int]complex128)
	values[0] = make(map[int]complex128)
	values[0][0] = 1
	values[0][1] = 0
	values[0][2] = 0
	vec = NewSparseMatrixFromMap(matrix.Cols(), 1, values)

	var lastEig complex128 = 0
	for {
		vec = vec.Scale(1 / cmplx.Sqrt(vec.Norm2Square()))
		vec = matrix.Dot(vec)
		eig = cmplx.Sqrt(vec.Norm2Square())
		if cmplx.Abs(lastEig-eig) < epsilon {
			break
		}
		lastEig = eig
	}

	return vec, eig
}
