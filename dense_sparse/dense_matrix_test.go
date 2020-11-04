package Singular

import (
	"fmt"
	"testing"
)

func TestDenseMatrix(t *testing.T) {
	eyes3 := NewDenseMatrixEye(3)
	fmt.Println(eyes3)
	fmt.Println(eyes3.Norm2Square())

	row1 := eyes3.GetRow(1)
	fmt.Println(row1)

	row1.Set(0, 1, 3)
	fmt.Println(eyes3)

	eyes3.SetRow(0, row1)
	fmt.Println(eyes3)

	col1 := eyes3.GetCol(1)
	fmt.Println(col1)

	col1.Set(2, 0, 10)
	eyes3.SetCol(1, col1)
	fmt.Println(eyes3)
}
