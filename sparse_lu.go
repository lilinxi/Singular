package Singular

func LU(A SparseMatrix) (L, U SparseMatrix) {
	if !A.IsSquare() {
		panic("")
	}
	U = SparseMatrixPrototype.New(A.Rows(), A.Cols())
	L = SparseMatrixPrototype.New(A.Rows(), A.Cols())
	U.SetRow(0, A.GetRow(0))
	L.SetCol(0, A.GetCol(0).Scale(1/U.Get(0, 0)))
	var lSlice, uSlice SparseMatrix
	for col := 1; col < A.Rows(); col++ {
		for row := col; row < A.Rows(); row++ {
			lSlice = L.GetSlice(col, col+1, 0, col)
			uSlice = U.GetSlice(0, col, row, row+1)
			U.Set(col, row, A.Get(col, row)-lSlice.Dot(uSlice).Get(0, 0))

			lSlice = L.GetSlice(row, row+1, 0, col)
			uSlice = U.GetSlice(0, col, col, col+1)
			L.Set(row, col, A.Get(row, col)-lSlice.Dot(uSlice).Get(0, 0))
			L.SetScale(row, col, 1/U.Get(col, col))
		}
	}
	return L, U
}

func LUSolve(A, b SparseMatrix) SparseMatrix {
	L, U := LU(A)

	// 求解 Ly = b
	y := b.Like()
	for row := 0; row < y.Rows(); row++ {
		value := b.Get(row, 0)
		for i := 0; i < row; i++ {
			value -= L.Get(row, i) * y.Get(i, 0)
		}
		y.Set(row, 0, value/L.Get(row, row))
	}

	// 求解 Ux = y
	x := b.Like()
	for row := y.Rows() - 1; row >= 0; row-- {
		value := y.Get(row, 0)
		for i := row+1; i < y.Rows(); i++ {
			value -= U.Get(row, i) * x.Get(i, 0)
		}
		x.Set(row, 0, value/U.Get(row, row))
	}

	return x
}

func InverseL(L SparseMatrix) SparseMatrix {
	matrix := L.Copy()
	inverse := SparseMatrixPrototype.Eyes(matrix.Rows())
	for i := 0; i < matrix.Rows()-1; i++ {
		inverse.SetRow(i, inverse.GetRow(i).Scale(matrix.Get(i, i)))
		matrix.SetRow(i, matrix.GetRow(i).Scale(matrix.Get(i, i)))
		for j := i; j < matrix.Rows(); j++ {
			inverse.SetRow(j,
				inverse.GetRow(j).Add(
					inverse.GetRow(i).Scale(-matrix.Get(j, i)),
				))
			matrix.SetRow(j,
				matrix.GetRow(j).Add(
					matrix.GetRow(i).Scale(-matrix.Get(j, i)),
				))
		}
	}
	inverse.SetRow(matrix.Rows()-1, inverse.GetRow(matrix.Rows()-1).Scale(matrix.Get(matrix.Rows()-1, matrix.Rows()-1)))
	matrix.SetRow(matrix.Rows()-1, matrix.GetRow(matrix.Rows()-1).Scale(matrix.Get(matrix.Rows()-1, matrix.Rows()-1)))
	return inverse
}

//// TODO: Error
//func InverseU(matrixU SparseMatrix) SparseMatrix {
//	matrixL := matrixU.Transpose()
//	fmt.Println("debug 1:", matrixL)
//	inverseL := InverseL(matrixL)
//	fmt.Println("debug 2:", inverseL)
//	fmt.Println("debug 3:", inverseL.Dot(matrixL))
//	fmt.Println("debug 4:", matrixL.Dot(inverseL))
//	inverseU := inverseL.Transpose()
//	return inverseU
//}
