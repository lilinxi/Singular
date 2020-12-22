package Singular

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func TestPowerMethod(t *testing.T) {
	A := SparseMatrixPrototype.From2DTable([][]float64{
		{1, 2, 3},
		{2, 1, 3},
		{3, 3, 5},
	})
	fmt.Println(A)
	//ans =
	//
	//-1.00000
	//-0.35890
	//8.35890

	eig, vec := SparsePowerMethod(A, 0, Epsilon)

	assert.T(t, Equal(eig, 8.358898943540675))
	assert.T(t, EqualSparseMatrix(vec, SparseMatrixPrototype.From1DList(
		[]float64{
			0.4389146460234735,
			0.4389146460234735,
			0.784033077753852},
	)))

	eig, vec = SparseInversePowerMethod(A, 0, Epsilon)

	assert.T(t, Equal(eig, -0.3588989435406737))
	assert.T(t, EqualSparseMatrix(vec, SparseMatrixPrototype.From1DList(
		[]float64{
			-0.5543951055176068,
			-0.5543951055176067,
			0.6207190459106626},
	)))
}
