package Singular

import (
	"bytes"
	"fmt"
)

type DenseMatrix struct {
	rows, cols int
	values     [][]complex128 // [rows][cols]complex128
}

func NewDenseMatrix(rows, cols int, values [][]complex128) DenseMatrix {
	return DenseMatrix{
		rows:   rows,
		cols:   cols,
		values: values,
	}
}

func NewDenseMatrixCopy(denseMatrix DenseMatrix) DenseMatrix {
	return DenseMatrix{
		rows:   denseMatrix.Rows(),
		cols:   denseMatrix.Cols(),
		values: copyValues2DTable(denseMatrix.values),
	}
}

func NewDenseMatrixFromSparseMatrix(rows, cols int, matrix SparseMatrix) DenseMatrix {
	values := full2DTable(rows, cols, 0)

	for row, rowValues := range matrix.values {
		for col, _ := range rowValues {
			values[row][col] = rowValues[col]
		}
	}

	return DenseMatrix{
		rows:   rows,
		cols:   cols,
		values: values,
	}
}

func full2DTable(rows, cols int, value complex128) [][]complex128 {
	cp := make([][]complex128, rows)
	for i := 0; i < rows; i++ {
		cp[i] = make([]complex128, cols)
		for j := 0; j < cols; j++ {
			cp[i][j] = value
		}
	}
	return cp
}

// 递归拷贝
func copyValues2DTable(values [][]complex128) [][]complex128 {
	cp := make([][]complex128, len(values))
	for i := 0; i < len(values); i++ {
		cp[i] = make([]complex128, len(values[0]))
		for j := 0; j < len(values[0]); j++ {
			cp[i][j] = values[i][j]
		}
	}
	return cp
}

func (m DenseMatrix) Rows() int { return m.rows }
func (m DenseMatrix) Cols() int { return m.cols }

func (m DenseMatrix) checkRowsCols(rows, cols int) {
	if rows >= m.Rows() || cols >= m.Cols() || rows < 0 || cols < 0 {
		panic("")
	}
}

func (m DenseMatrix) Get(rows, cols int) complex128 {
	m.checkRowsCols(rows, cols)

	return m.values[rows][cols]
}

func (m *DenseMatrix) Set(rows, cols int, value complex128) {
	m.checkRowsCols(rows, cols)
	m.values[rows][cols] = value
}

func (m *DenseMatrix) SetAdd(rows, cols int, value complex128) {
	m.checkRowsCols(rows, cols)
	m.values[rows][cols] += value
}

func (m *DenseMatrix) SetMul(rows, cols int, value complex128) {
	m.checkRowsCols(rows, cols)
	m.values[rows][cols] *= value
}

func (m DenseMatrix) Add(matrix DenseMatrix) DenseMatrix {
	if m.Rows() != matrix.Rows() || m.Cols() != matrix.Cols() {
		panic("")
	}

	retValues := copyValues2DTable(m.values)

	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			retValues[i][j] += matrix.values[i][j]
		}
	}

	return DenseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: retValues,
	}
}

func (m DenseMatrix) Sub(matrix DenseMatrix) DenseMatrix {
	if m.Rows() != matrix.Rows() || m.Cols() != matrix.Cols() {
		panic("")
	}

	retValues := copyValues2DTable(m.values)

	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			retValues[i][j] -= matrix.values[i][j]
		}
	}

	return DenseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: retValues,
	}
}

func (m DenseMatrix) Dot(matrix DenseMatrix) DenseMatrix {
	if m.Cols() != matrix.Rows() {
		panic("")
	}

	retValues := full2DTable(m.Rows(), matrix.Cols(), 0)

	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < matrix.Cols(); col++ {
			for k := 0; k < m.Cols(); k++ {
				retValues[row][col] += m.values[row][k] * matrix.values[k][col]
			}
		}
	}

	return DenseMatrix{
		rows:   m.rows,
		cols:   matrix.cols,
		values: retValues,
	}
}

func (m DenseMatrix) Scale(scale complex128) DenseMatrix {
	retValues := copyValues2DTable(m.values)

	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			retValues[i][j] *= scale
		}
	}

	return DenseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: retValues,
	}
}

func (m DenseMatrix) Transpose() DenseMatrix {
	retValues := make([][]complex128, m.Cols())
	for i := 0; i < m.Cols(); i++ {
		retValues[i] = make([]complex128, m.Rows())
		for j := 0; j < m.Rows(); j++ {
			retValues[i][j] = m.values[j][i]
		}
	}

	return DenseMatrix{
		rows:   m.cols,
		cols:   m.rows,
		values: retValues,
	}
}

func (m DenseMatrix) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%dX%d:\n", m.Rows(), m.Cols()))

	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			buf.WriteString(fmt.Sprintf("%v", m.Get(i, j)))
			buf.WriteString(" ")
		}
		buf.WriteString("\b")
		buf.WriteString("\n")
	}

	return buf.String()
}
