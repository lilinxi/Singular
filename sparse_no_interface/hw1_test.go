package Singular

import (
	"fmt"
	"math"
	"testing"
)

const N = 501

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

	eig1, _ := PowerMethodDev2(A, 0, Epsilon)
	fmt.Println(eig1)

	eig2, _ := InversePowerMethod(A, 0, Epsilon)
	fmt.Println(eig2)
}

func TestHw1_2(t *testing.T) {
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

	eig1, vec := PowerMethod(A, 0, Epsilon)
	if ErrorSparseMatrix(vec.Scale(eig1), A.Dot(vec)) > Epsilon {
		fmt.Println("eig1:", eig1, "error:", ErrorSparseMatrix(vec.Scale(eig1), A.Dot(vec)))
		eig1, vec = PowerMethod(A, eig1, Epsilon)
	}
	fmt.Println("eig1:", eig1, "error:", ErrorSparseMatrix(vec.Scale(eig1), A.Dot(vec)))

	eig2, _ := InversePowerMethod(A, 0, Epsilon)
	if ErrorSparseMatrix(vec.Scale(eig2), A.Dot(vec)) > Epsilon {
		fmt.Println("eig2:", eig2, "error:", ErrorSparseMatrix(vec.Scale(eig2), A.Dot(vec)))
		eig2, vec = InversePowerMethod(A, eig2, Epsilon)
	}
	fmt.Println("eig2:", eig2, "error:", ErrorSparseMatrix(vec.Scale(eig2), A.Dot(vec)))
}
