package Singular

import (
	"fmt"
	"math"
	"testing"
)

func TestHw1(t *testing.T) {
	A := NewSparseMatrix(501, 501)
	a := func(i int) float64 {
		fi := float64(i)
		return (1.64-0.024*fi)*math.Sin(0.2*fi) - 0.64*math.Exp(0.1/2)
	}
	b := 0.16
	c := -0.064

	for i := 0; i < 501; i++ {
		A.Set(i, i, a(i))
	}

	for i := 1; i < 501; i++ {
		A.Set(i, i-1, b)
		A.Set(i-1, i, b)
	}

	for i := 2; i < 501; i++ {
		A.Set(i, i-2, c)
		A.Set(i-2, i, c)
	}

	fmt.Println(A)

	eig, _ := PowerMethod(A, Epsilon)
	fmt.Println(eig)
	//fmt.Println(vec)

	l,_:=LU(NewDenseMatrixFromSparseMatrix(501,501, A))
	fmt.Println(l)
}

func TestHw1_1(t *testing.T) {
	A := NewDenseMatrix(501, 501)
	a := func(i int) float64 {
		fi := float64(i)
		return (1.64-0.024*fi)*math.Sin(0.2*fi) - 0.64*math.Exp(0.1/2)
	}
	b := 0.16
	c := -0.064

	for i := 0; i < 501; i++ {
		A.Set(i, i, a(i))
	}

	for i := 1; i < 501; i++ {
		A.Set(i, i-1, b)
		A.Set(i-1, i, b)
	}

	for i := 2; i < 501; i++ {
		A.Set(i, i-2, c)
		A.Set(i-2, i, c)
	}

	fmt.Println(A)

	l,u:=LU(A)
	fmt.Println(l)
	fmt.Println(u)
	fmt.Println(l.Dot(u).Add(A.Scale(-1)).Norm2Square())
}
