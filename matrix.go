package Singular

type MatrixInterface interface {
	Cols() int
	Rows() int

	New(cols, rows int) MatrixInterface
	Get(cols, rows int) complex64
	Set(cols, rows int, value complex64)

	Add(matrix MatrixInterface) MatrixInterface
	Sub(matrix MatrixInterface) MatrixInterface
	Dot(matrix MatrixInterface) MatrixInterface
	Scale(scale complex64) MatrixInterface

	Transpose() MatrixInterface

	String() string
}

type DefaultMatrix struct{}

func (DefaultMatrix) Cols() int { panic("") }
func (DefaultMatrix) Rows() int { panic("") }

func (DefaultMatrix) New(cols, rows int) MatrixInterface  { panic("") }
func (DefaultMatrix) Get(cols, rows int) complex64        { panic("") }
func (DefaultMatrix) Set(cols, rows int, value complex64) { panic("") }

func (DefaultMatrix) Add(matrix MatrixInterface) MatrixInterface { panic("") }
func (DefaultMatrix) Sub(matrix MatrixInterface) MatrixInterface { panic("") }
func (DefaultMatrix) Dot(matrix MatrixInterface) MatrixInterface {
	//var mat [4][4]float64
	//for i := 0; i < 4; i++ {
	//	for j := 0; j < 4; j++ {
	//		mat[i][j] =
	//			m1.mat[i][0]*m2.mat[0][j] +
	//				m1.mat[i][1]*m2.mat[1][j] +
	//				m1.mat[i][2]*m2.mat[2][j] +
	//				m1.mat[i][3]*m2.mat[3][j]
	//	}
	//}
	//return Matrix4x4{mat: mat}
	panic("")
}
func (DefaultMatrix) Scale(scale complex64) MatrixInterface { panic("") }

func (DefaultMatrix) Transpose() MatrixInterface { panic("") }

func (DefaultMatrix) String() string { panic("") }
