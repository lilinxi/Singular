package Singular

import (
	"bytes"
	"fmt"
	"math"
)

var DenseMatrixPrototype = DenseMatrix{
	rows:   0,
	cols:   0,
	values: nil,
}

type DenseMatrix struct {
	rows, cols int
	values     [][]float64 // [rows][cols]float64
}

// 递归拷贝
func copyDenseValues(values [][]float64) [][]float64 {
	cp := make([][]float64, len(values))
	for row := 0; row < len(values); row++ {
		cp[row] = make([]float64, len(values[0]))
		for col := 0; col < len(values[0]); col++ {
			cp[row][col] = values[row][col]
		}
	}
	return cp
}

func fullDenseValues(rows, cols int, value float64) [][]float64 {
	cp := make([][]float64, rows)
	for row := 0; row < rows; row++ {
		cp[row] = make([]float64, cols)
		for col := 0; col < cols; col++ {
			cp[row][col] = value
		}
	}
	return cp
}

func (m DenseMatrix) Rows() int         { return m.rows }
func (m DenseMatrix) Cols() int         { return m.cols }
func (m DenseMatrix) IsRowVector() bool { return m.Rows() == 1 }
func (m DenseMatrix) IsColVector() bool { return m.Cols() == 1 }
func (m DenseMatrix) IsSquare() bool    { return m.Rows() == m.Cols() }

func (m DenseMatrix) assertRowsCols(rows, cols int) {
	if rows >= m.Rows() || cols >= m.Cols() || rows < 0 || cols < 0 {
		panic(fmt.Sprintf("error: (%d, %d) must limit: (%d, %d)", rows, cols, m.Rows(), m.Cols()))
	}
}

func (m DenseMatrix) assertShapeMatch(matrix DenseMatrix) {
	if m.Rows() != matrix.Rows() || m.Cols() != matrix.Cols() {
		panic(fmt.Sprintf("error: (%d, %d) must match: (%d, %d)", matrix.Rows(), matrix.Cols(), m.Rows(), m.Cols()))
	}
}

func (m DenseMatrix) assertDotMatch(matrix DenseMatrix) {
	if m.Cols() != matrix.Rows() {
		panic(fmt.Sprintf("error: (%d, %d) must match: (%d, %d)", matrix.Rows(), matrix.Cols(), m.Rows(), m.Cols()))
	}
}

func (m DenseMatrix) Get(rows, cols int) float64 {
	m.assertRowsCols(rows, cols)

	return m.values[rows][cols]
}

func (m *DenseMatrix) Set(rows, cols int, value float64) {
	m.assertRowsCols(rows, cols)

	m.values[rows][cols] = value
}

func (m *DenseMatrix) SetAdd(rows, cols int, value float64) {
	m.assertRowsCols(rows, cols)

	m.values[rows][cols] += value
}

func (m *DenseMatrix) SetScale(rows, cols int, value float64) {
	m.assertRowsCols(rows, cols)

	m.values[rows][cols] *= value
}

func (m DenseMatrix) Add(matrix DenseMatrix) DenseMatrix {
	m.assertShapeMatch(matrix)

	cpValues := copyDenseValues(m.values)

	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			cpValues[row][col] += matrix.values[row][col]
		}
	}

	return DenseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: cpValues,
	}
}

func (m DenseMatrix) Sub(matrix DenseMatrix) DenseMatrix {
	return m.Add(matrix.Scale(-1))
}

func (m DenseMatrix) Scale(scale float64) DenseMatrix {
	cpValues := copyDenseValues(m.values)

	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			cpValues[row][col] *= scale
		}
	}

	return DenseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: cpValues,
	}
}

func (m DenseMatrix) Dot(matrix DenseMatrix) DenseMatrix {
	m.assertDotMatch(matrix)

	retValues := fullDenseValues(m.Rows(), matrix.Cols(), 0)

	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < matrix.Cols(); col++ {
			for k := 0; k < m.Cols(); k++ {
				retValues[row][col] += m.values[row][k] * matrix.values[k][col]
			}
		}
	}

	return DenseMatrix{
		rows:   m.Rows(),
		cols:   matrix.Cols(),
		values: retValues,
	}
}

func (m DenseMatrix) Transpose() DenseMatrix {
	retMatrix := m.Zeros(m.Cols(), m.Rows())

	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			retMatrix.Set(col, row, m.Get(row, col))
		}
	}

	return retMatrix
}

func (m DenseMatrix) NormK(k float64) float64 {
	if k == 1 {
		var norm float64
		for row := 0; row < m.Rows(); row++ {
			for col := 0; col < m.Cols(); col++ {
				norm += math.Abs(m.Get(row, col))
			}
		}
		return norm
	}
	var norm float64
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			norm += math.Pow(m.Get(row, col), k)
		}
	}
	return math.Pow(norm, 1/k)
}

func (m DenseMatrix) NormK2Square() float64 {
	var norm2square float64
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			value := m.Get(row, col)
			norm2square += value * value
		}
	}
	return norm2square
}

func (m DenseMatrix) NormK2() float64 {
	return math.Sqrt(m.NormK2Square())
}

func (m DenseMatrix) NormInf() float64 {
	var normInf float64
	for row := 0; row < m.Rows(); row++ {
		normInf = math.Max(normInf, m.GetRow(row).NormK(1))
	}
	return normInf
}

func (m DenseMatrix) Normal() DenseMatrix {
	norm := math.Sqrt(m.NormK2Square())
	retMatrix := m.Copy().Scale(1 / norm)
	return retMatrix
}

func (m DenseMatrix) GetSlice(rowBegin, rowEnd, colBegin, colEnd int) DenseMatrix {
	m.assertRowsCols(rowBegin, colBegin)
	m.assertRowsCols(rowEnd-1, colEnd-1)

	rows := rowEnd - rowBegin
	cols := colEnd - colBegin

	if rows <= 0 || cols <= 0 {
		panic(fmt.Sprintf("slice(rows, cols): (%d, %d)", rows, cols))
	}

	var slice = m.Zeros(rows, cols)

	for row := 0; row < slice.Rows(); row++ {
		for col := 0; col < slice.Cols(); col++ {
			slice.Set(row, col, m.Get(row+rowBegin, col+colBegin))
		}
	}
	return slice
}

func (m *DenseMatrix) SetSlice(rowBegin, colBegin int, matrix DenseMatrix) {
	rowEnd := rowBegin + matrix.Rows()
	colEnd := colBegin + matrix.Cols()

	m.assertRowsCols(rowBegin, colBegin)
	m.assertRowsCols(rowEnd-1, colEnd-1)

	rows := rowEnd - rowBegin
	cols := colEnd - colBegin

	if rows <= 0 || cols <= 0 {
		panic(fmt.Sprintf("slice(rows, cols): (%d, %d)", rows, cols))
	}

	for row := rowBegin; row < rowEnd; row++ {
		for col := colBegin; col < colEnd; col++ {
			m.Set(row, col, matrix.Get(row-rowBegin, col-colBegin))
		}
	}
}

func (m DenseMatrix) GetRow(row int) DenseMatrix {
	m.assertRowsCols(row, 0)
	return m.GetSlice(row, row+1, 0, m.Cols())
}

func (m DenseMatrix) GetCol(col int) DenseMatrix {
	m.assertRowsCols(0, col)
	return m.GetSlice(0, m.Rows(), col, col+1)
}

func (m *DenseMatrix) SetRow(row int, matrix DenseMatrix) {
	m.assertRowsCols(row, 0)

	if !matrix.IsRowVector() || m.Cols() != matrix.Cols() {
		panic(fmt.Sprintf("error: (%d, %d) must limit: (%d, %d)", 1, m.Cols(), m.Rows(), m.Cols()))
	}

	for j := 0; j < m.Cols(); j++ {
		m.Set(row, j, matrix.Get(0, j))
	}
}

func (m *DenseMatrix) SetCol(col int, matrix DenseMatrix) {
	m.assertRowsCols(0, col)

	if !matrix.IsColVector() || m.Rows() != matrix.Rows() {
		panic(fmt.Sprintf("error: (%d, %d) must limit: (%d, %d)", m.Rows(), 1, m.Rows(), m.Cols()))
	}

	for i := 0; i < m.Rows(); i++ {
		m.Set(i, col, matrix.Get(i, 0))
	}
}

func (m DenseMatrix) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%dX%d:\n", m.Rows(), m.Cols()))

	buf.WriteString("[\n")
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			buf.WriteString(fmt.Sprintf("%v", m.Get(row, col)))
			buf.WriteString(" ")
		}
		buf.WriteString("\b;")
		buf.WriteString("\n")
	}
	buf.WriteString("]")

	return buf.String()
}

func (m DenseMatrix) Copy() DenseMatrix {
	return DenseMatrix{
		rows:   m.Rows(),
		cols:   m.Cols(),
		values: copyDenseValues(m.values),
	}
}

func (m DenseMatrix) Like() DenseMatrix {
	return DenseMatrix{
		rows:   m.Rows(),
		cols:   m.Cols(),
		values: fullDenseValues(m.Rows(), m.Cols(), 0),
	}
}

func (m DenseMatrix) New(rows, cols int) DenseMatrix {
	return DenseMatrix{
		rows:   rows,
		cols:   cols,
		values: fullDenseValues(rows, cols, 0),
	}
}

func (m DenseMatrix) Zeros(rows, cols int) DenseMatrix {
	return DenseMatrix{
		rows:   rows,
		cols:   cols,
		values: fullDenseValues(rows, cols, 0),
	}
}

func (m DenseMatrix) Eyes(size int) DenseMatrix {
	retMatrix := m.Zeros(size, size)

	for i := 0; i < size; i++ {
		retMatrix.Set(i, i, 1)
	}

	return retMatrix
}

func (m DenseMatrix) Full(rows, cols int, value float64) DenseMatrix {
	return DenseMatrix{
		rows:   rows,
		cols:   cols,
		values: fullDenseValues(rows, cols, value),
	}
}

func (m DenseMatrix) From2DTable(valueTable [][]float64) DenseMatrix {
	return DenseMatrix{
		rows:   len(valueTable),
		cols:   len(valueTable[0]),
		values: valueTable,
	}
}

func (m DenseMatrix) From1DList(valueList []float64) DenseMatrix {
	retMatrix := m.Zeros(len(valueList), 1)

	for row := 0; row < retMatrix.Rows(); row++ {
		retMatrix.Set(row, 0, valueList[row])
	}

	return retMatrix
}

/**
由四个分块矩阵构造矩阵
*/
func (m DenseMatrix) FromBlocks(leftTop, rightTop, leftBottom, rightBottom DenseMatrix) DenseMatrix {
	// 分块矩阵的各个维度应该相匹配
	if leftTop.Rows() != rightTop.Rows() ||
		leftTop.Cols() != leftBottom.Cols() ||
		rightTop.Cols() != rightBottom.Cols() ||
		leftBottom.Rows() != rightBottom.Rows() {
		panic("blocks mis match")
	}

	rows := leftTop.Rows() + leftBottom.Rows()
	cols := leftTop.Cols() + rightTop.Cols()

	retMatrix := DenseMatrixPrototype.New(rows, cols)

	retMatrix.SetSlice(0, 0, leftTop)
	retMatrix.SetSlice(leftTop.Rows(), 0, leftBottom)
	retMatrix.SetSlice(0, leftTop.Cols(), rightTop)
	retMatrix.SetSlice(leftTop.Rows(), leftTop.Cols(), rightBottom)

	return retMatrix
}