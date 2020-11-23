package Singular

import (
	"fmt"
	"math"
)

func PowerMethodDev(matrix SparseMatrix, offset, epsilon float64) (eig float64, vec SparseMatrix) {
	if !matrix.Square() {
		panic("")
	}

	if offset != 0 {
		matrix = matrix.Add(NewSparseMatrixEyes(matrix.Rows()).Scale(offset))
	}

	vec = NewSparseMatrixFull(matrix.Rows(), 1, 1)
	vec.Scale(vec.NormK(2))

	var lastEig float64 = 0
	vec = matrix.Dot(vec)
	lastEig = vec.NormK(2)
	vec = vec.Scale(1 / lastEig)
	for {
		vecTmp := matrix.Dot(vec)
		eig = vecTmp.NormK(2)
		vecTmp = vecTmp.Scale(1 / eig)
		err := ErrorSparseMatrix(vec, vecTmp)
		vec = vecTmp
		fmt.Println("err:", err)
		if err < epsilon {
			break
		}
		lastEig = eig
	}

	if vec.Get(0, 0)*matrix.Dot(vec).Get(0, 0) < 0 {
		eig = -eig
	}

	return eig - offset, vec
}

func PowerMethodDev2(matrix SparseMatrix, offset, epsilon float64) (eig float64, vec SparseMatrix) {
	if !matrix.Square() {
		panic("")
	}

	if offset != 0 {
		matrix = matrix.Add(NewSparseMatrixEyes(matrix.Rows()).Scale(offset))
	}

	vec = NewSparseMatrixFull(matrix.Rows(), 1, 1)
	vec.Scale(vec.NormK(1))

	var lastEig float64 = 0
	vec = matrix.Dot(vec)
	lastEig = vec.NormK(1)
	vec = vec.Scale(1 / lastEig)
	for {
		vec = matrix.Dot(vec)
		eig = vec.NormK(1)
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

func PowerMethod(matrix SparseMatrix, offset, epsilon float64) (eig float64, vec SparseMatrix) {
	if !matrix.Square() {
		panic("")
	}

	if offset != 0 {
		matrix = matrix.Add(NewSparseMatrixEyes(matrix.Rows()).Scale(offset))
	}

	vec = NewSparseMatrixFull(matrix.Rows(), 1, 1)
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
