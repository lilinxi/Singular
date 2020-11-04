package Singular

import (
	"bytes"
	"fmt"
)

type DenseMatrix struct {
	rows, cols int
	values     [][]float64 // [rows][cols]float64
}

func NewDenseMatrix(rows, cols int) DenseMatrix {
	return NewDenseMatrixFrom2DTable(rows, cols, full2DTable(rows, cols, 0))
}

func NewDenseMatrixOnes(size int) DenseMatrix {
	return NewDenseMatrixFrom2DTable(size, size, full2DTable(size, size, 1))
}

func NewDenseMatrixZeros(size int) DenseMatrix {
	return NewDenseMatrixFrom2DTable(size, size, full2DTable(size, size, 0))
}

func NewDenseMatrixEye(size int) DenseMatrix {
	values := full2DTable(size, size, 0)
	for i := 0; i < size; i++ {
		values[i][i] = 1
	}
	return NewDenseMatrixFrom2DTable(size, size, values)
}

func NewDenseMatrixFrom2DTable(rows, cols int, values [][]float64) DenseMatrix {
	return DenseMatrix{
		rows:   rows,
		cols:   cols,
		values: values,
	}
}

func NewDenseMatrixFromCopy(denseMatrix DenseMatrix) DenseMatrix {
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

func full2DTable(rows, cols int, value float64) [][]float64 {
	cp := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		cp[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			cp[i][j] = value
		}
	}
	return cp
}

// 递归拷贝
func copyValues2DTable(values [][]float64) [][]float64 {
	cp := make([][]float64, len(values))
	for i := 0; i < len(values); i++ {
		cp[i] = make([]float64, len(values[0]))
		for j := 0; j < len(values[0]); j++ {
			cp[i][j] = values[i][j]
		}
	}
	return cp
}

func (m DenseMatrix) Rows() int { return m.rows }
func (m DenseMatrix) Cols() int { return m.cols }

func (m DenseMatrix) Square() bool {
	return m.Rows() == m.Cols()
}

func (m DenseMatrix) checkRowsCols(rows, cols int) {
	if rows >= m.Rows() || cols >= m.Cols() || rows < 0 || cols < 0 {
		panic("")
	}
}

func (m DenseMatrix) Get(rows, cols int) float64 {
	m.checkRowsCols(rows, cols)

	return m.values[rows][cols]
}

func (m DenseMatrix) GetRow(rows int) DenseMatrix {
	m.checkRowsCols(rows, 0)

	rowValue := make([][]float64, 1)
	rowValue[0] = make([]float64, m.Cols())
	for i := 0; i < m.Cols(); i++ {
		rowValue[0][i] = m.Get(rows, i)
	}

	return DenseMatrix{
		rows:   1,
		cols:   m.Cols(),
		values: rowValue,
	}
}

func (m DenseMatrix) GetCol(cols int) DenseMatrix {
	m.checkRowsCols(0, cols)

	colValue := make([][]float64, m.Rows())
	for i := 0; i < m.Cols(); i++ {
		colValue[i] = make([]float64, 1)
		colValue[i][0] = m.Get(i, cols)
	}

	return DenseMatrix{
		rows:   m.Rows(),
		cols:   1,
		values: colValue,
	}
}

func (m *DenseMatrix) Set(rows, cols int, value float64) {
	m.checkRowsCols(rows, cols)
	m.values[rows][cols] = value
}

func (m *DenseMatrix) SetRow(rows int, rowValue DenseMatrix) {
	m.checkRowsCols(rows, 0)
	for i := 0; i < m.Cols(); i++ {
		m.Set(rows, i, rowValue.Get(0, i))
	}
}

func (m *DenseMatrix) SetCol(cols int, colValue DenseMatrix) {
	m.checkRowsCols(0, cols)
	for i := 0; i < m.Rows(); i++ {
		m.Set(i, cols, colValue.Get(i, 0))
	}
}

func (m *DenseMatrix) SetAdd(rows, cols int, value float64) {
	m.checkRowsCols(rows, cols)
	m.values[rows][cols] += value
}

func (m *DenseMatrix) SetScale(rows, cols int, value float64) {
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

func (m DenseMatrix) Scale(scale float64) DenseMatrix {
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
	retValues := make([][]float64, m.Cols())
	for i := 0; i < m.Cols(); i++ {
		retValues[i] = make([]float64, m.Rows())
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

func (m DenseMatrix) Norm2Square() float64 {
	var norm2square float64
	for _, values := range m.values {
		for _, value := range values {
			norm2square += value * value
		}
	}
	return norm2square
}

func (m DenseMatrix) Slice(axis, begin, end int) DenseMatrix {
	var slice DenseMatrix
	if axis == 0 {
		if begin < 0 || begin > m.Rows() || end < 0 || end > m.Rows()+1 {
			panic("")
		}
		rows := end - begin
		slice = NewDenseMatrix(rows, m.Cols())
		for i := 0; i < rows; i++ {
			for j := 0; j < m.Cols(); j++ {
				slice.Set(i, j, m.Get(i+begin, j))
			}
		}
	} else if axis == 1 {
		if begin < 0 || begin > m.Cols() || end < 0 || end > m.Cols()+1 {
			panic("")
		}
		cols := end - begin
		slice = NewDenseMatrix(m.Rows(), cols)
		for i := 0; i < m.Rows(); i++ {
			for j := 0; j < cols; j++ {
				slice.Set(i, j, m.Get(i+begin, j))
			}
		}
	} else {
		panic("")
	}
	return slice
}

func (m DenseMatrix) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%dX%d:\n", m.Rows(), m.Cols()))

	buf.WriteString("[\n")
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			buf.WriteString(fmt.Sprintf("%v", m.Get(i, j)))
			buf.WriteString(" ")
		}
		buf.WriteString("\b;")
		buf.WriteString("\n")
	}
	buf.WriteString("]")

	return buf.String()
}
