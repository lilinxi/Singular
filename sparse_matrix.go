package Singular

import (
	"bytes"
	"fmt"
)

const Epsilon = 1e-12

type SparseMatrix struct {
	rows, cols int
	values     map[int]map[int]complex64 // map[rows]map[cols]complex64
}

type Tuple struct {
	row, col int
	value    complex64
}

func NewSparseMatrix(rows, cols int, values map[int]map[int]complex64) SparseMatrix {
	return SparseMatrix{
		rows:   rows,
		cols:   cols,
		values: values,
	}
}

func NewSparseMatrixCopy(sparseMatrix SparseMatrix) SparseMatrix {
	return SparseMatrix{
		rows:   sparseMatrix.Rows(),
		cols:   sparseMatrix.Cols(),
		values: CopyValues(sparseMatrix.values),
	}
}

func NewSparseMatrixFromTupleList(rows, cols int, tupleList []Tuple) SparseMatrix {
	values := make(map[int]map[int]complex64)

	for _, tuple := range tupleList {
		if _, ok := values[tuple.row]; !ok {
			values[tuple.row] = make(map[int]complex64)
		}
		values[tuple.row][tuple.col] = tuple.value
	}

	return SparseMatrix{
		rows:   rows,
		cols:   cols,
		values: values,
	}
}

func NewSparseMatrixFrom2DTable(rows, cols int, valueTable [][]complex64) SparseMatrix {
	values := make(map[int]map[int]complex64)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if valueTable[row][col] != 0 {
				if _, ok := values[row]; !ok {
					values[row] = make(map[int]complex64)
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

// 递归拷贝
func CopyValues(values map[int]map[int]complex64) map[int]map[int]complex64 {
	cp := make(map[int]map[int]complex64)
	for row, values := range values {
		cp[row] = make(map[int]complex64)
		for col, value := range values {
			cp[row][col] = value
		}
	}
	return cp
}

func (m SparseMatrix) Rows() int { return m.rows }
func (m SparseMatrix) Cols() int { return m.cols }

func (m SparseMatrix) Get(rows, cols int) complex64 {
	if rows >= m.Rows() || cols >= m.Cols() || rows < 0 || cols < 0 {
		panic("")
	}

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

func (m *SparseMatrix) Set(rows, cols int, value complex64) {
	if _, ok := m.values[rows]; !ok {
		m.values[rows] = make(map[int]complex64)
	}
	m.values[rows][cols] = value
}

func (m *SparseMatrix) SetAdd(rows, cols int, value complex64) {
	if _, ok := m.values[rows]; !ok {
		m.values[rows] = make(map[int]complex64)
	}
	if _, ok := m.values[rows][cols]; !ok {
		m.values[rows][cols] = 0
	}
	m.values[rows][cols] += value
}

func (m *SparseMatrix) SetMul(rows, cols int, value complex64) {
	if _, ok := m.values[rows]; !ok {
		m.values[rows] = make(map[int]complex64)
	}
	if _, ok := m.values[rows][cols]; !ok {
		return
	}
	m.values[rows][cols] *= value
}

func (m SparseMatrix) Add(matrix SparseMatrix) SparseMatrix {
	if m.Rows() != matrix.Rows() || m.Cols() != matrix.Cols() {
		panic("")
	}

	retValues := CopyValues(m.values)

	for row, values := range matrix.values {
		if _, ok := retValues[row]; !ok {
			retValues[row] = make(map[int]complex64)
		}
		for col, value := range values {
			if _, ok := retValues[row][col]; !ok {
				retValues[row][col] = value
			} else {
				retValues[row][col] += value
			}
		}
	}

	return SparseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: retValues,
	}
}

func (m SparseMatrix) Sub(matrix SparseMatrix) SparseMatrix {
	if m.Rows() != matrix.Rows() || m.Cols() != matrix.Cols() {
		panic("")
	}

	retValues := CopyValues(m.values)

	for row, values := range matrix.values {
		if _, ok := retValues[row]; !ok {
			retValues[row] = make(map[int]complex64)
		}
		for col, value := range values {
			if _, ok := retValues[row][col]; !ok {
				retValues[row][col] = -value
			} else {
				retValues[row][col] -= value
			}
		}
	}

	return SparseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: retValues,
	}
}

func (m SparseMatrix) Dot(matrix SparseMatrix) SparseMatrix {
	if m.Cols() != matrix.Rows() {
		panic("")
	}

	retValues := make(map[int]map[int]complex64)

	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < matrix.Cols(); col++ {
			var value complex64 = 0
			for k := 0; k < m.Cols(); k++ {
				if m.values[row][k] == 0 || matrix.values[k][col] == 0 {
					continue
				}
				value += m.values[row][k] * matrix.values[k][col]
			}
			if value != 0 {
				if _, ok := retValues[row]; !ok {
					retValues[row] = make(map[int]complex64)
				}
				retValues[row][col] = value
			}
		}
	}

	return SparseMatrix{
		rows:   m.rows,
		cols:   matrix.cols,
		values: retValues,
	}
}

func (m SparseMatrix) Scale(scale complex64) SparseMatrix {
	retValues := CopyValues(m.values)

	for row, values := range m.values {
		for col, _ := range values {
			retValues[row][col] *= scale
		}
	}

	return SparseMatrix{
		rows:   m.rows,
		cols:   m.cols,
		values: retValues,
	}
}

func (m SparseMatrix) Transpose() SparseMatrix {
	retValues := make(map[int]map[int]complex64)

	for row, values := range m.values {
		for col, value := range values {
			if _, ok := retValues[col]; !ok {
				retValues[col] = make(map[int]complex64)
			}
			retValues[col][row] = value
		}
	}

	return SparseMatrix{
		rows:   m.cols,
		cols:   m.rows,
		values: retValues,
	}
}

func (m SparseMatrix) String() string {
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
