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

	//fmt.Println(A)

	eig1, _ := PowerMethod(A, 0, Epsilon)
	fmt.Println(eig1)

	eig2, _ := InversePowerMethod(A, 0, Epsilon)
	fmt.Println(eig2)

	//for i := 10.0; i > -10; i -= 1 {
	//	Acpy := A.Add(NewSparseMatrixEyes(N).Scale(i))
	//	eig, _ := PowerMethod(Acpy, 0, Epsilon)
	//	fmt.Println(eig, eig-i)
	//}

	//fmt.Println(vec)

	//l,_:=LU(A)
	//fmt.Println(l)
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

	eig1, _ := PowerMethod(A, 0, Epsilon)
	fmt.Println(eig1)

	eig1, _ = PowerMethod(A, eig1, Epsilon)
	fmt.Println(eig1)

	eig2, _ := InversePowerMethod(A, 0, Epsilon)
	fmt.Println(eig2)
	//
	//eig3 := eig2 - eig1
	//
	//lamdba1 := math.Min(eig1, eig3)
	//lamdba2 := math.Max(eig1, eig3)
	//fmt.Println(lamdba1, lamdba2)

}
