package Singular

//import (
//	"bytes"
//	"fmt"
//)
//
///**
//go 的一个不好用的地方就是没有办法在父类中调用子类的方法，所以就写不出来模板方法，只能使用自下而上的继承，或者使用自下而上的代理
//*/
//
//type Matrixable interface {
//	SetData(data Matrixable) // 指针接收器
//
//	Rows() int
//	Cols() int
//
//	IsRowVector() bool
//	IsColVector() bool
//	IsSquare() bool
//
//	checkRowCol(rows, cols int)
//
//	Get(rows, cols int) float64
//	Set(rows, cols int, value float64)      // 指针接收器
//	SetAdd(rows, cols int, value float64)   // 指针接收器
//	SetScale(rows, cols int, value float64) // 指针接收器
//
//	Add(matrix Matrix) Matrix
//	Scale(scale float64) Matrix
//	Dot(matrix Matrix) Matrix
//	Transpose() Matrix
//
//	Norm2Square() float64
//	Norm(pow float64) float64
//	NormInf() float64
//
//	GetSlice(rowBegin, rowEnd, colBegin, colEnd int) Matrix
//	SetSlice(rowBegin, colBegin int, matrix Matrix) // 指针接收器
//	GetRow(row int) Matrix
//	GetCol(col int) Matrix
//	SetRow(row int, matrix Matrix) // 指针接收器
//	SetCol(col int, matrix Matrix) // 指针接收器
//
//	String() string
//
//	Copy() Matrix
//	Like() Matrix
//	//Zeros() Matrix
//	//Eyes() Matrixable
//	//Full() Matrixable
//	//From2DTable(rows, cols int, valueTable [][]float64)
//	//From1DList(valueList []float64)
//}
//
//type MatrixDatable interface {
//	Matrixable
//	SetMatrixAgent(matrix Matrix)
//}
//
//type MatrixData struct {
//	Matrix
//}
//
//
//
//
//type Matrix struct {
//	rows, cols int
//	data       MatrixData
//}
//
//func NewMatrix(rows, cols int) Matrix {
//	return Matrix{
//		rows: rows,
//		cols: cols,
//	}
//}
//
//func NewMatrixAgent
//
//// 指针接收器
//func (m *Matrix) SetData(data Matrixable) {
//	m.data = data
//}
//
//func (m Matrix) Rows() int { return m.rows }
//func (m Matrix) Cols() int { return m.cols }
//
//func (m Matrix) IsRowVector() bool { return m.Rows() == 1 }
//func (m Matrix) IsColVector() bool { return m.Cols() == 1 }
//func (m Matrix) IsSquare() bool    { return m.Rows() == m.Cols() }
//
//func (m Matrix) checkRowCol(rows, cols int) {
//	if rows >= m.Rows() || cols >= m.Cols() || rows < 0 || cols < 0 {
//		panic(fmt.Sprintf("param: %d, %d must limit in %d, %d", rows, cols, m.Rows(), m.Cols()))
//	}
//}
//
//func (m Matrix) Get(rows, cols int) float64              { panic("no impl") }
//func (m *Matrix) Set(rows, cols int, value float64)      { panic("no impl") } // 指针接收器
//func (m *Matrix) SetAdd(rows, cols int, value float64)   { panic("no impl") } // 指针接收器
//func (m *Matrix) SetScale(rows, cols int, value float64) { panic("no impl") } // 指针接收器
//
//func (m Matrix) Add(matrix Matrix) Matrix   { panic("no impl") }
//func (m Matrix) Scale(scale float64) Matrix { panic("no impl") }
//func (m Matrix) Dot(matrix Matrix) Matrix   { panic("no impl") }
//func (m Matrix) Transpose() Matrix          { panic("no impl") }
//
//func (m Matrix) Norm2Square() float64     { panic("no impl") }
//func (m Matrix) Norm(pow float64) float64 { panic("no impl") }
//func (m Matrix) NormInf() float64         { panic("no impl") }
//
//func (m Matrix) GetSlice(rowBegin, rowEnd, colBegin, colEnd int) Matrix { panic("no impl") }
//func (m *Matrix) SetSlice(rowBegin, colBegin int, matrix Matrix)        { panic("no impl") } // 指针接收器
//func (m Matrix) GetRow(row int) Matrix                                  { panic("no impl") }
//func (m Matrix) GetCol(col int) Matrix                                  { panic("no impl") }
//func (m *Matrix) SetRow(row int, matrix Matrix)                         { panic("no impl") } // 指针接收器
//func (m *Matrix) SetCol(col int, matrix Matrix)                         { panic("no impl") } // 指针接收器
//
//func (m Matrix) String() string {
//	var buf bytes.Buffer
//
//	buf.WriteString(fmt.Sprintf("%dX%d:\n", m.Rows(), m.Cols()))
//
//	buf.WriteString("[\n")
//	for i := 0; i < m.Rows(); i++ {
//		for j := 0; j < m.Cols(); j++ {
//			buf.WriteString(fmt.Sprintf("%v", m.data.Get(i, j)))
//			buf.WriteString(" ")
//		}
//		buf.WriteString("\b;")
//		buf.WriteString("\n")
//	}
//	buf.WriteString("]")
//
//	return buf.String()
//}
//
//func (m Matrix) Copy() Matrix { panic("no impl") }
//func (m Matrix) Like() Matrix { panic("no impl") }
