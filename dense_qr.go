package Singular

import (
	"fmt"
	"sort"
)

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

/**
QR 迭代法计算所有特征值
Ak = Qk·Rk
Ak+1 = Rk·Qk
Ak 与 Ak+1 相似，且 Ak+1 收敛到对角阵，其对角线即为特征值
*/
func DenseQRIter(matrix DenseMatrix) DenseMatrix {
	if !matrix.IsSquare() {
		panic("matrix must be square")
	}
	n := matrix.Rows()
	A := matrix.Copy()
	for true {
		Q, R := DenseQR(A)
		Atmp := R.Dot(Q)
		if EqualDenseMatrix(A, Atmp) {
			break
		}
		A = Atmp
	}
	lambda := make([]float64, n)
	for i := 0; i < n; i++ {
		lambda[i] = A.Get(i, i)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(lambda))) // 降序排序，[Reverse 包装了原 Less 方法](https://itimetraveler.github.io/2016/09/07/%E3%80%90Go%E8%AF%AD%E8%A8%80%E3%80%91%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B%E6%8E%92%E5%BA%8F%E5%92%8C%20slice%20%E6%8E%92%E5%BA%8F/)
	return DenseMatrixPrototype.From1DList(lambda)
}

/**
带双步位移的 QR 迭代法计算所有特征值，收敛速度更快
Ak-sk·I = Qk·Rk
Ak+1 = Rk·Qk+sk·I
Ak 与 Ak+1 相似，且 Ak+1 收敛到对角阵，其对角线即为特征值
sk = a_nn^k
*/
func DenseQRDisplaceIter(matrix DenseMatrix) DenseMatrix {
	if !matrix.IsSquare() {
		panic("matrix must be square")
	}
	n := matrix.Rows()
	A := matrix.Copy()
	for true {
		sk := A.Get(n-1, n-1)
		Q, R := DenseQR(A.Sub(DenseMatrixPrototype.Eyes(n).Scale(sk)))
		Atmp := R.Dot(Q).Add(DenseMatrixPrototype.Eyes(n).Scale(sk))
		if EqualDenseMatrix(A, Atmp) {
			break
		}
		fmt.Println(Atmp)
		A = Atmp
	}
	lambda := make([]float64, n)
	for i := 0; i < n; i++ {
		lambda[i] = A.Get(i, i)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(lambda))) // 降序排序，[Reverse 包装了原 Less 方法](https://itimetraveler.github.io/2016/09/07/%E3%80%90Go%E8%AF%AD%E8%A8%80%E3%80%91%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B%E6%8E%92%E5%BA%8F%E5%92%8C%20slice%20%E6%8E%92%E5%BA%8F/)
	return DenseMatrixPrototype.From1DList(lambda)
}
