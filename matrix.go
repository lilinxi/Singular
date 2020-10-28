package Singular

type MatrixInterface interface {
	Rows() int
	Cols() int

	Get(cols, rows int) complex128
	Set(cols, rows int, value complex128)
}

type DefaultMatrix struct {
	rows, cols int
}

func (m DefaultMatrix) Rows() int { return m.rows }
func (m DefaultMatrix) Cols() int { return m.cols }

func (m DefaultMatrix) Get(cols, rows int) complex128        { panic("") }
func (m DefaultMatrix) Set(cols, rows int, value complex128) { panic("") }
