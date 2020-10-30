package Singular

func LU(A DenseMatrix) (L, U DenseMatrix) {
	if !A.Square() {
		panic("")
	}
	U = NewDenseMatrixZeros(A.Rows())
	L = NewDenseMatrixZeros(A.Rows())
	U.SetRow(0, A.GetRow(0))
	L.SetCol(0, A.GetCol(0).Scale(1/U.Get(0, 0)))
	var lSlice, uSlice DenseMatrix
	for col := 1; col < A.Rows(); col++ {
		for row := col; row < A.Rows(); row++ {
			lSlice = L.GetRow(col).Slice(1, 0, col)
			uSlice = U.GetCol(row).Slice(0, 0, col)
			U.Set(col, row, A.Get(col, row)-lSlice.Dot(uSlice).Get(0, 0))

			lSlice = L.GetRow(row).Slice(1, 0, col)
			uSlice = U.GetCol(col).Slice(0, 0, col)
			L.Set(row, col, A.Get(row, col)-lSlice.Dot(uSlice).Get(0, 0))
			L.SetScale(row, col, 1/U.Get(col, col))
		}
	}
	return L, U
}
