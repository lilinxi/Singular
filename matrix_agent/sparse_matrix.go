package Singular

//type SparseMatrix struct {
//	MatrixData
//	values map[int]map[int]float64 // map[rows]map[cols]float64
//}

//func NewSparseMatrixFrom2DTable(rows, cols int, valueTable [][]float64) SparseMatrix {
//	values := make(map[int]map[int]float64)
//
//	for row := 0; row < rows; row++ {
//		for col := 0; col < cols; col++ {
//			if valueTable[row][col] != 0 {
//				if _, ok := values[row]; !ok {
//					values[row] = make(map[int]float64)
//				}
//				values[row][col] = valueTable[row][col]
//			}
//		}
//	}
//
//	return SparseMatrix{
//		Matrix: NewMatrix(rows, cols),
//		values: values,
//	}
//}

//func NewSparseMatrixFrom1DList(valueList []float64) Matrix {
//	values := make(map[int]map[int]float64)
//
//	for col, value := range valueList {
//		if value == 0 {
//			continue
//		} else {
//			values[col] = make(map[int]float64)
//			values[col][0] = value
//		}
//	}
//
//	matrix := NewMatrix(len(valueList), 1)
//	data := &SparseMatrix{
//		Matrix: matrix,
//		values: values,
//	}
//	matrix.SetData(data)
//
//	return matrix
//}

//type Tuple struct {
//	row, col int
//	value    float64
//}
//
//func NewSparseMatrix(rows, cols int) SparseMatrix {
//	return SparseMatrix{
//		rows:   rows,
//		cols:   cols,
//		values: make(map[int]map[int]float64),
//	}
//}
//
//func NewSparseMatrixZeros(size int) SparseMatrix {
//	return SparseMatrix{
//		rows:   size,
//		cols:   size,
//		values: make(map[int]map[int]float64),
//	}
//}
//
//func NewSparseMatrixEyes(size int) SparseMatrix {
//	ret := SparseMatrix{
//		rows:   size,
//		cols:   size,
//		values: make(map[int]map[int]float64),
//	}
//	for i := 0; i < size; i++ {
//		ret.Set(i, i, 1)
//	}
//	return ret
//}
//
//func NewSparseMatrixFull(rows, cols int, value float64) SparseMatrix {
//	ret := SparseMatrix{
//		rows:   rows,
//		cols:   cols,
//		values: make(map[int]map[int]float64),
//	}
//	for i := 0; i < rows; i++ {
//		for j := 0; j < cols; j++ {
//			ret.Set(i, j, value)
//		}
//	}
//	return ret
//}
//
//func NewSparseMatrixFromMap(rows, cols int, values map[int]map[int]float64) SparseMatrix {
//	return SparseMatrix{
//		rows:   rows,
//		cols:   cols,
//		values: values,
//	}
//}
//
//func NewSparseMatrixCopy(sparseMatrix SparseMatrix) SparseMatrix {
//	return SparseMatrix{
//		rows:   sparseMatrix.Rows(),
//		cols:   sparseMatrix.Cols(),
//		values: copyValues(sparseMatrix.values),
//	}
//}
//

//}
//
//func NewSparseMatrixFrom1DListd(valueList []float64) *SparseMatrix {
//	values := make(map[int]map[int]float64)
//
//	for col, value := range valueList {
//		if value == 0 {
//			continue
//		} else {
//			values[col] = make(map[int]float64)
//			values[col][0] = value
//		}
//	}
//
//	return &SparseMatrix{
//		rows:   len(valueList),
//		cols:   1,
//		values: values,
//	}
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
//
//// 递归拷贝
//func copyValues(values map[int]map[int]float64) map[int]map[int]float64 {
//	cp := make(map[int]map[int]float64)
//	for row, values := range values {
//		cp[row] = make(map[int]float64)
//		for col, value := range values {
//			cp[row][col] = value
//		}
//	}
//	return cp
//}
//
//func (m *SparseMatrix) Rows() int        { return m.rows }
//func (m *SparseMatrix) Cols() int        { return m.cols }
//func (m SparseMatrix) IsRowVector() bool { return m.Rows() == 1 }
//func (m SparseMatrix) IsColVector() bool { return m.Cols() == 1 }
//func (m SparseMatrix) IsSquare() bool    { return m.Rows() == m.Cols() }
//
//func (m SparseMatrix) checkRowCol(rows, cols int) {
//	if rows >= m.Rows() || cols >= m.Cols() || rows < 0 || cols < 0 {
//		panic(fmt.Sprintf("param: %d, %d must limit in %d, %d", rows, cols, m.Rows(), m.Cols()))
//	}
//}
//
//func (m SparseMatrix) Get(rows, cols int) float64 {
//	if rows >= m.Rows() || cols >= m.Cols() || rows < 0 || cols < 0 {
//		panic(fmt.Sprintf("error: %d, %d, limit: %d, %d", rows, cols, m.Rows(), m.Cols()))
//	}
//
//	values, ok := m.values[rows]
//	if !ok {
//		return 0
//	}
//	value, ok := values[cols]
//	if !ok {
//		return 0
//	}
//	return value
//}

//
//func (m *SparseMatrix) Set(rows, cols int, value float64) {
//	if value == 0 {
//		return
//	}
//	if _, ok := m.values[rows]; !ok {
//		m.values[rows] = make(map[int]float64)
//	}
//	m.values[rows][cols] = value
//}
//
//func (m *SparseMatrix) SetAdd(rows, cols int, value float64) {
//	if _, ok := m.values[rows]; !ok {
//		m.values[rows] = make(map[int]float64)
//	}
//	if _, ok := m.values[rows][cols]; !ok {
//		m.values[rows][cols] = 0
//	}
//	m.values[rows][cols] += value
//}
//
//func (m *SparseMatrix) SetScale(rows, cols int, value float64) {
//	if _, ok := m.values[rows]; !ok {
//		m.values[rows] = make(map[int]float64)
//	}
//	if _, ok := m.values[rows][cols]; !ok {
//		return
//	}
//	m.values[rows][cols] *= value
//}
//
//func (m SparseMatrix) Add(matrix SparseMatrix) SparseMatrix {
//	if m.Rows() != matrix.Rows() || m.Cols() != matrix.Cols() {
//		panic("")
//	}
//
//	retValues := copyValues(m.values)
//
//	for row, values := range matrix.values {
//		if _, ok := retValues[row]; !ok {
//			retValues[row] = make(map[int]float64)
//		}
//		for col, value := range values {
//			if _, ok := retValues[row][col]; !ok {
//				retValues[row][col] = value
//			} else {
//				retValues[row][col] += value
//			}
//		}
//	}
//
//	return SparseMatrix{
//		rows:   m.rows,
//		cols:   m.cols,
//		values: retValues,
//	}
//}
//
//func (m SparseMatrix) Scale(scale float64) SparseMatrix {
//	retValues := copyValues(m.values)
//
//	for row, values := range m.values {
//		for col, _ := range values {
//			retValues[row][col] *= scale
//		}
//	}
//
//	return SparseMatrix{
//		rows:   m.rows,
//		cols:   m.cols,
//		values: retValues,
//	}
//}
//
//func (m SparseMatrix) Dot(matrix SparseMatrix) SparseMatrix {
//	if m.Cols() != matrix.Rows() {
//		panic("")
//	}
//
//	retValues := make(map[int]map[int]float64)
//
//	for row := 0; row < m.Rows(); row++ {
//		for col := 0; col < matrix.Cols(); col++ {
//			var value float64 = 0
//			for k := 0; k < m.Cols(); k++ {
//				if m.values[row][k] == 0 || matrix.values[k][col] == 0 {
//					continue
//				}
//				value += m.values[row][k] * matrix.values[k][col]
//			}
//			if value != 0 {
//				if _, ok := retValues[row]; !ok {
//					retValues[row] = make(map[int]float64)
//				}
//				retValues[row][col] = value
//			}
//		}
//	}
//
//	return SparseMatrix{
//		rows:   m.rows,
//		cols:   matrix.cols,
//		values: retValues,
//	}
//}
//
//func (m SparseMatrix) Transpose() SparseMatrix {
//	retValues := make(map[int]map[int]float64)
//
//	for row, values := range m.values {
//		for col, value := range values {
//			if _, ok := retValues[col]; !ok {
//				retValues[col] = make(map[int]float64)
//			}
//			retValues[col][row] = value
//		}
//	}
//
//	return SparseMatrix{
//		rows:   m.cols,
//		cols:   m.rows,
//		values: retValues,
//	}
//}
//
//func (m SparseMatrix) Norm2Square() float64 {
//	var norm2square float64
//	for _, values := range m.values {
//		for _, value := range values {
//			norm2square += value * value
//		}
//	}
//	return norm2square
//}
//
//func (m SparseMatrix) NormK(pow float64) float64 {
//	if pow == 1 {
//		var norm float64
//		for _, values := range m.values {
//			for _, value := range values {
//				norm += math.Abs(value)
//			}
//		}
//		return norm
//	}
//	var norm float64
//	for _, values := range m.values {
//		for _, value := range values {
//			norm += math.Pow(value, pow)
//		}
//	}
//	return math.Pow(norm, 1/pow)
//}
//
//func (m SparseMatrix) NormInf() float64 {
//	var normInf float64
//	for i := 0; i < m.Rows(); i++ {
//		normInf = math.Max(normInf, m.GetRow(i).NormK(1))
//	}
//	return normInf
//}
//
//func (m SparseMatrix) GetSlice(rowBegin, rowEnd, colBegin, colEnd int) SparseMatrix {
//	m.checkRowCol(rowBegin, colBegin)
//	m.checkRowCol(rowEnd-1, colEnd-1)
//	row := rowEnd - rowBegin
//	col := colEnd - colBegin
//	if row <= 0 || col <= 0 {
//		panic("")
//	}
//	var slice = NewSparseMatrix(row, col)
//	for i := 0; i < slice.Rows(); i++ {
//		for j := 0; j < slice.Cols(); j++ {
//			slice.Set(i, j, m.Get(i+rowBegin, j+colBegin))
//		}
//	}
//	return slice
//}
//
//func (m *SparseMatrix) SetSlice(rowBegin, colBegin int, matrix SparseMatrix) {
//	rowEnd := rowBegin + matrix.Rows()
//	colEnd := colBegin + matrix.Cols()
//
//	m.checkRowCol(rowBegin, colBegin)
//	m.checkRowCol(rowEnd-1, colEnd-1)
//
//	row := rowEnd - rowBegin
//	col := colEnd - colBegin
//	if row <= 0 || col <= 0 {
//		panic("")
//	}
//
//	for i := rowBegin; i < rowEnd; i++ {
//		for j := colBegin; j < colEnd; j++ {
//			m.Set(i, j, matrix.Get(i-rowBegin, j-colBegin))
//		}
//	}
//}
//
//func (m SparseMatrix) GetRow(row int) SparseMatrix {
//	m.checkRowCol(row, 0)
//	return m.GetSlice(row, row+1, 0, m.Cols())
//}
//
//func (m SparseMatrix) GetCol(col int) SparseMatrix {
//	m.checkRowCol(0, col)
//	return m.GetSlice(0, m.Rows(), col, col+1)
//}
//
//func (m *SparseMatrix) SetRow(row int, matrix SparseMatrix) {
//	m.checkRowCol(row, 0)
//
//	if !matrix.IsRowVector() || m.Cols() != matrix.Cols() {
//		panic("")
//	}
//
//	for j := 0; j < m.Cols(); j++ {
//		m.Set(row, j, matrix.Get(0, j))
//	}
//}
//
//func (m *SparseMatrix) SetCol(col int, matrix SparseMatrix) {
//	m.checkRowCol(0, col)
//
//	if !matrix.IsColVector() || m.Rows() != matrix.Rows() {
//		panic("")
//	}
//
//	for i := 0; i < m.Rows(); i++ {
//		m.Set(i, col, matrix.Get(i, 0))
//	}
//}
//
//func (m SparseMatrix) String() string {
//	var buf bytes.Buffer
//
//	buf.WriteString(fmt.Sprintf("%dX%d:\n", m.Rows(), m.Cols()))
//
//	buf.WriteString("[\n")
//	for i := 0; i < m.Rows(); i++ {
//		for j := 0; j < m.Cols(); j++ {
//			buf.WriteString(fmt.Sprintf("%v", m.Get(i, j)))
//			buf.WriteString(" ")
//		}
//		buf.WriteString("\b;")
//		buf.WriteString("\n")
//	}
//	buf.WriteString("]")
//
//	return buf.String()
//}

///**
//复制矩阵
//*/
//func (m SparseMatrix) Copy() Matrix {
//	return SparseMatrix{
//		rows:   m.Rows(),
//		cols:   m.Cols(),
//		values: copyValues(m.values),
//	}
//}
//
///**
//复制矩阵的大小
//*/
//func (m SparseMatrix) Like() Matrix {
//	return SparseMatrix{
//		rows: m.cols,
//		cols: m.Rows(),
//	}
//}
