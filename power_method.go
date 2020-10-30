package Singular

import (
	"fmt"
	"math"
)

func PowerMethod(matrix SparseMatrix, epsilon float64) (eig float64, vec SparseMatrix) {
	values := make(map[int]map[int]float64)
	values[0] = make(map[int]float64)
	values[0][0] = 1
	values[0][1] = 1
	values[0][2] = 1
	vec = NewSparseMatrixFromMap(matrix.Cols(), 1, values)

	var lastEig float64 = 0
	for {
		vec = vec.Scale(1 / math.Sqrt(vec.Norm2Square()))
		vec = matrix.Dot(vec)
		eig = math.Sqrt(vec.Norm2Square())
		if math.Abs(lastEig-eig) < epsilon {
			break
		}
		lastEig = eig
		fmt.Println("eig: ", eig)
	}

	return eig, vec
}
