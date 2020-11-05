package Singular

import (
	"math"
)

func PowerMethodDev(matrix SparseMatrix, epsilon float64) (eig float64, vec SparseMatrix) {
	vec = NewSparseMatrixFrom1DList([]float64{1, 1, 1})

	var lastEig float64 = 0
	for {
		vec = vec.Scale(1 / math.Sqrt(vec.Norm2Square()))
		vec = matrix.Dot(vec)
		eig = math.Sqrt(vec.Norm2Square())
		if math.Abs(lastEig-eig) < epsilon {
			break
		}
		lastEig = eig
		//fmt.Println("eig: ", eig)
	}

	return eig, vec
}

func PowerMethod(matrix SparseMatrix, offset, epsilon float64) (eig float64, vec SparseMatrix) {
	if !matrix.Square() {
		panic("")
	}

	if offset != 0 {
		matrix = matrix.Add(NewSparseMatrixEyes(matrix.Rows()).Scale(offset))
	}

	vec = NewSparseMatrixFull(matrix.Rows(), 1, 1)
	vec.Scale(vec.Norm(2))

	var lastEig float64 = 0
	vec = matrix.Dot(vec)
	lastEig = vec.Norm(2)
	vec = vec.Scale(1 / lastEig)
	for {
		vec = matrix.Dot(vec)
		eig = vec.Norm(2)
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
