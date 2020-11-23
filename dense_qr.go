package Singular

/**
由向量 w 构建 Householder 矩阵，为反射矩阵，对称正交矩阵
*/
func DenseHouseholder(w DenseMatrix) DenseMatrix {
	if !Equal(w.NormK2Square(), 1) {
		panic("w must be normal")
	}
	wt := w.Transpose()
	I := DenseMatrixPrototype.Eyes(wt.Cols())
	H := I.Add(w.Dot(wt).Scale(-2))
	return H
}

/**
QR 分解将矩阵分解为一个正交阵 Q 和一个上三角阵 R，其中 Householder 矩阵必为正交阵，且可以构造 Householder 矩阵消去对角线一下的元素为 0
构造 Householder 矩阵，将向量反射到 (σ,0,0,0……)^T，σ 应该取和 xk 相反的符号，xk 是保持不为零的元素，这样做是为了避免精度丢失

*/
func DenseQR(matrix DenseMatrix) (Q, R DenseMatrix) {
	R = matrix.Copy()
	col := R.GetCol(0)
	newCol := col.Like()
	sigma := col.NormK2()
	// σ 应该取和 xk 相反的符号，xk 是保持不为零的元素，这样做是为了避免精度丢失
	if col.Get(0, 0) > 0 {
		sigma *= -1
	}
	newCol.Set(0, 0, sigma)
	w := col.Sub(newCol).Normal()
	H := DenseHouseholder(w)
	Q = H
	R = H.Dot(R)
	for i := 1; i < matrix.Cols()-1; i++ {
		colSlice := R.GetSlice(i, matrix.Rows(), i, i+1)
		newColSlice := colSlice.Like()
		sigma := colSlice.NormK2()
		// σ 应该取和 xk 相反的符号，xk 是保持不为零的元素，这样做是为了避免精度丢失
		if colSlice.Get(0, 0) > 0 {
			sigma *= -1
		}
		newColSlice.Set(0, 0, sigma)
		w := colSlice.Sub(newColSlice).Normal()
		H = DenseMatrixPrototype.FromBlocks(
			DenseMatrixPrototype.Eyes(i),
			DenseMatrixPrototype.Zeros(i, matrix.Cols()-i),
			DenseMatrixPrototype.Zeros(matrix.Rows()-i, i),
			DenseHouseholder(w))
		Q = Q.Dot(H)
		R = H.Dot(R)
	}
	return Q, R
}

func DenseQRInter(matrix DenseMatrix) DenseMatrix {
	A := matrix.Copy()
	for true {
		Q, R := DenseQR(A)
		Atmp := R.Dot(Q)
		if EqualDenseMatrix(A, Atmp) {
			return Atmp
		}
		A = Atmp
	}
	panic("unexpected error")
}
