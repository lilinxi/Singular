package Singular

/**
go 的一个不好用的地方就是没有办法在父类中调用子类的方法，所以就写不出来模板方法，只能使用自下而上的继承，或者使用自下而上的代理
这里使用原型模式，方便矩阵 Dense 和 Sparse 的区分
go 传参的指针和值没有太大的区别，只有在赋值到接口时能不能使用指针接收器的区别
*/

type Matrixable interface {
	// 属性
	Rows() int
	Cols() int

	// 属性判断
	IsRowVector() bool
	IsColVector() bool
	IsSquare() bool

	// 规则断言
	assertRowsCols(rows, cols int)
	assertShapeMatch(matrix Matrixable)
	assertDotMatch(matrix Matrixable)

	// 简单 getter-setter
	Get(rows, cols int) float64
	Set(rows, cols int, value float64)      // 指针接收器
	SetAdd(rows, cols int, value float64)   // 指针接收器
	SetScale(rows, cols int, value float64) // 指针接收器

	// 数学运算
	Add(matrix Matrixable) Matrixable
	Scale(scale float64) Matrixable
	Dot(matrix Matrixable) Matrixable
	Transpose() Matrixable

	// 范数
	NormK(k float64) float64 // k 阶范数
	NormK2Square() float64   // 2 阶范数的平方
	NormInf() float64        // 无穷范数

	// slice getter-setter
	GetSlice(rowBegin, rowEnd, colBegin, colEnd int) Matrixable
	SetSlice(rowBegin, colBegin int, matrix Matrixable) // 指针接收器
	GetRow(row int) Matrixable
	GetCol(col int) Matrixable
	SetRow(row int, matrix Matrixable) // 指针接收器
	SetCol(col int, matrix Matrixable) // 指针接收器

	// strinter
	String() string

	// 原型构造器
	Copy() Matrixable
	Like() Matrixable

	// 复制原型构造器
	Zeros(rows, cols int) Matrixable
	Eyes(size int) Matrixable
	Full(rows, cols int, value float64) Matrixable
	From2DTable(valueTable [][]float64) Matrixable
	From1DList(valueList []float64) Matrixable
}
