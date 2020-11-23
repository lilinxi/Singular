package Singular

import (
	"bytes"
	"fmt"
	"math"
)

var SparseMatrixPrototype = SparseMatrix{
	rows:   0,
	cols:   0,
	values: nil,
}

type SparseMatrix struct {
	rows, cols int
	values     map[int]map[int]float64 // map[rows]map[cols]float64
}

//func NewSparseMatrixFrom1DList(valueList []float64) SparseMatrix {

//}
//
//func NewSparseMatrixFromTupleList(rows, cols int, tupleList []Tuple) SparseMatrix {
//	values := make(map[int]map[int]float64)
//
//	for _, tuple := range tupleList {
//		if _, ok := values[tuple.row]; !ok {
//			values[tuple.row] = make(map[int]float64)
//		}
//		values[tuple.row][tuple.col] = tuple.value
//	}
//
//	return SparseMatrix{
//		rows:   rows,
//		cols:   cols,
//		values: values,
//	}
//}

// 递归拷贝
func copySparseValues(values map[int]map[int]float64) map[int]map[int]float64 {
	cp := make(map[int]map[int]float64)
	for row, values := range values {
		cp[row] = make(map[int]float64)
		for col, value := range values {
			cp[row][col] = value
		}
	}
	return cp
}

func (m SparseMatrix) Rows() int         { return m.rows }
func (m SparseMatrix) Cols() int         { return m.cols }
func (m SparseMatrix) IsRowVector() bool { return m.Rows() == 1 }
func (m SparseMatrix) IsColVector() bool { return m.Cols() == 1 }
func (m SparseMatrix) IsSquare() bool    { return m.Rows() == m.Cols() }

func (m SparseMatrix) assertRowsCols(rows, cols int) {
	if rows >= m.Rows() || cols >= m.Cols() || rows < 0 || cols < 0 {
		panic(fmt.Sprintf("error: (%d, %d) must limit: (%d, %d)", rows, cols, m.Rows(), m.Cols()))
	}
}

func (m SparseMatrix) assertShapeMatch(matrix SparseMatrix) {
	if m.Rows() != matrix.Rows() || m.Cols() != matrix.Cols() {
		panic(fmt.Sprintf("error: (%d, %d) must match: (%d, %d)", matrix.Rows(), matrix.Cols(), m.Rows(), m.Cols()))
	}
}

func (m SparseMatrix) assertDotMatch(matrix SparseMatrix) {
	if m.Cols() != matrix.Rows() {
		panic(fmt.Sprintf("error: (%d, %d) must match: (%d, %d)", matrix.Rows(), matrix.Cols(), m.Rows(), m.Cols()))
	}
}

func (m SparseMatrix) Get(rows, cols int) float64 {
	m.assertRowsCols(rows, cols)

	values, ok := m.values[rows]
	if !ok {
		return 0
	}
	value, ok := values[cols]
	if !ok {
		return 0
	}
	return value
}

func (m *SparseMatrix) Set(rows, cols int, value float64) {
	m.assertRowsCols(rows, cols)

	if value == 0 {
		return
	}
	if _, ok := m.values[rows]; !ok {
		m.values[rows] = make(map[int]float64)
	}
	m.values[rows][cols] = value
}

func (m *SparseMatrix) SetAdd(rows, cols int, value float64) {
	m.assertRowsCols(rows, cols)

	if _, ok := m.values[rows]; !ok {
		m.values[rows] = make(map[int]float64)
	}
	if _, ok := m.values[rows][cols]; !ok {
		m.values[rows][cols] = 0
	}
	m.values[rows][cols] += value
}

func (m *SparseMatrix) SetScale(rows, cols int, value float64) {
	m.assertRowsCols(rows, cols)

	if _, ok := m.values[rows]; !ok {
		m.values[rows] = make(map[int]float64)
	}
	if _, ok := m.values[rows][cols]; !ok {
		return
	}
	m.values[rows][cols] *= value
}

func (m SparseMatrix) Add(matrix SparseMatrix) SparseMatrix {
	m.assertShapeMatch(matrix)

	cpValues := copySparseValues(m.values)

	for row, values := range matrix.values {
		if _, ok := cpValues[row]; !ok {
			cpValues[row] = make(map[int]float64)
		}
		for col, value := range values {
			if _, ok := cpValues[row][col]; !ok {
				cpValues[row][col] = value
			} else {
				cpValues[row][col] += value
			}
		}
	}

	return SparseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: cpValues,
	}
}

func (m SparseMatrix) Sub(matrix SparseMatrix) SparseMatrix {
	return m.Add(matrix.Scale(-1))
}

func (m SparseMatrix) Scale(scale float64) SparseMatrix {
	cpValues := copySparseValues(m.values)

	for row, values := range m.values {
		for col, _ := range values {
			cpValues[row][col] *= scale
		}
	}

	return SparseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: cpValues,
	}
}

func (m SparseMatrix) Dot(matrix SparseMatrix) SparseMatrix {
	m.assertDotMatch(matrix)

	cpValues := make(map[int]map[int]float64)

	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < matrix.Cols(); col++ {
			var value float64 = 0
			for k := 0; k < m.Cols(); k++ {
				if m.values[row][k] == 0 || matrix.values[k][col] == 0 {
					continue
				}
				value += m.values[row][k] * matrix.values[k][col]
			}
			if value != 0 {
				if _, ok := cpValues[row]; !ok {
					cpValues[row] = make(map[int]float64)
				}
				cpValues[row][col] = value
			}
		}
	}

	return SparseMatrix{
		rows:   m.rows,
		cols:   matrix.cols,
		values: cpValues,
	}
}

func (m SparseMatrix) Transpose() SparseMatrix {
	cpValues := make(map[int]map[int]float64)

	for row, values := range m.values {
		for col, value := range values {
			if _, ok := cpValues[col]; !ok {
				cpValues[col] = make(map[int]float64)
			}
			cpValues[col][row] = value
		}
	}

	return SparseMatrix{
		rows:   m.cols,
		cols:   m.rows,
		values: cpValues,
	}
}

func (m SparseMatrix) NormK(pow float64) float64 {
	if pow == 1 {
		var norm float64
		for _, values := range m.values {
			for _, value := range values {
				norm += math.Abs(value)
			}
		}
		return norm
	}
	var norm float64
	for _, values := range m.values {
		for _, value := range values {
			norm += math.Pow(value, pow)
		}
	}
	return math.Pow(norm, 1/pow)
}

func (m SparseMatrix) NormK2Square() float64 {
	var norm2square float64
	for _, values := range m.values {
		for _, value := range values {
			norm2square += value * value
		}
	}
	return norm2square
}

func (m SparseMatrix) NormK2() float64 {
	return math.Sqrt(m.NormK2Square())
}

func (m SparseMatrix) NormInf() float64 {
	var normInf float64
	for i := 0; i < m.Rows(); i++ {
		normInf = math.Max(normInf, m.GetRow(i).NormK(1))
	}
	return normInf
}

func (m SparseMatrix) Normal() SparseMatrix {
	norm := math.Sqrt(m.NormK2Square())
	retMatrix := m.Copy().Scale(1 / norm)
	return retMatrix
}

func (m SparseMatrix) GetSlice(rowBegin, rowEnd, colBegin, colEnd int) SparseMatrix {
	m.assertRowsCols(rowBegin, colBegin)
	m.assertRowsCols(rowEnd-1, colEnd-1)

	row := rowEnd - rowBegin
	col := colEnd - colBegin

	if row <= 0 || col <= 0 {
		panic("")
	}

	var slice = SparseMatrixPrototype.New(row, col)

	for i := 0; i < slice.Rows(); i++ {
		for j := 0; j < slice.Cols(); j++ {
			slice.Set(i, j, m.Get(i+rowBegin, j+colBegin))
		}
	}
	return slice
}

func (m *SparseMatrix) SetSlice(rowBegin, colBegin int, matrix SparseMatrix) {
	rowEnd := rowBegin + matrix.Rows()
	colEnd := colBegin + matrix.Cols()

	m.assertRowsCols(rowBegin, colBegin)
	m.assertRowsCols(rowEnd-1, colEnd-1)

	row := rowEnd - rowBegin
	col := colEnd - colBegin
	if row <= 0 || col <= 0 {
		panic("")
	}

	for i := rowBegin; i < rowEnd; i++ {
		for j := colBegin; j < colEnd; j++ {
			m.Set(i, j, matrix.Get(i-rowBegin, j-colBegin))
		}
	}
}

func (m SparseMatrix) GetRow(row int) SparseMatrix {
	m.assertRowsCols(row, 0)
	return m.GetSlice(row, row+1, 0, m.Cols())
}

func (m SparseMatrix) GetCol(col int) SparseMatrix {
	m.assertRowsCols(0, col)
	return m.GetSlice(0, m.Rows(), col, col+1)
}

func (m *SparseMatrix) SetRow(row int, matrix SparseMatrix) {
	m.assertRowsCols(row, 0)

	if !matrix.IsRowVector() || m.Cols() != matrix.Cols() {
		panic("")
	}

	for j := 0; j < m.Cols(); j++ {
		m.Set(row, j, matrix.Get(0, j))
	}
}

func (m *SparseMatrix) SetCol(col int, matrix SparseMatrix) {
	m.assertRowsCols(0, col)

	if !matrix.IsColVector() || m.Rows() != matrix.Rows() {
		panic("")
	}

	for i := 0; i < m.Rows(); i++ {
		m.Set(i, col, matrix.Get(i, 0))
	}
}

func (m SparseMatrix) String() string {
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

func (m SparseMatrix) Copy() SparseMatrix {
	return SparseMatrix{
		rows:   m.Rows(),
		cols:   m.Cols(),
		values: copySparseValues(m.values),
	}
}

func (m SparseMatrix) Like() SparseMatrix {
	return SparseMatrix{
		rows:   m.Rows(),
		cols:   m.Cols(),
		values: make(map[int]map[int]float64),
	}
}

func (m SparseMatrix) New(rows, cols int) SparseMatrix {
	return SparseMatrix{
		rows:   rows,
		cols:   cols,
		values: make(map[int]map[int]float64),
	}
}

func (m SparseMatrix) Zeros(rows, cols int) SparseMatrix {
	return SparseMatrix{
		rows:   rows,
		cols:   cols,
		values: make(map[int]map[int]float64),
	}
}

func (m SparseMatrix) Eyes(size int) SparseMatrix {
	retMatrix := m.Zeros(size, size)

	for i := 0; i < size; i++ {
		retMatrix.Set(i, i, 1)
	}

	return retMatrix
}

func (m SparseMatrix) Full(rows, cols int, value float64) SparseMatrix {
	panic("using dense matrix instead")
}

func (m SparseMatrix) From2DTable(valueTable [][]float64) SparseMatrix {
	values := make(map[int]map[int]float64)
	rows := len(valueTable)
	cols := len(valueTable[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if valueTable[row][col] != 0 {
				if _, ok := values[row]; !ok {
					values[row] = make(map[int]float64)
				}
				values[row][col] = valueTable[row][col]
			}
		}
	}

	return SparseMatrix{
		rows:   rows,
		cols:   cols,
		values: values,
	}
}

func (m SparseMatrix) From1DList(valueList []float64) SparseMatrix {
	values := make(map[int]map[int]float64)

	for col, value := range valueList {
		if value == 0 {
			continue
		} else {
			values[col] = make(map[int]float64)
			values[col][0] = value
		}
	}

	return SparseMatrix{
		rows:   len(valueList),
		cols:   1,
		values: values,
	}
}

/**
由四个分块矩阵构造矩阵
*/
func (m SparseMatrix) FromBlocks(leftTop, rightTop, leftBottom, rightBottom SparseMatrix) SparseMatrix {
	// 分块矩阵的各个维度应该相匹配
	if leftTop.Rows() != rightTop.Rows() ||
		leftTop.Cols() != leftBottom.Cols() ||
		rightTop.Cols() != rightBottom.Cols() ||
		leftBottom.Rows() != rightBottom.Rows() {
		panic("blocks mis match")
	}

	rows := leftTop.Rows() + leftBottom.Rows()
	cols := leftTop.Cols() + rightTop.Cols()

	retMatrix := SparseMatrixPrototype.New(rows, cols)

	retMatrix.SetSlice(0, 0, leftTop)
	retMatrix.SetSlice(leftTop.Rows(), 0, leftBottom)
	retMatrix.SetSlice(0, leftTop.Cols(), rightTop)
	retMatrix.SetSlice(leftTop.Rows(), leftTop.Cols(), rightBottom)

	return retMatrix
}
