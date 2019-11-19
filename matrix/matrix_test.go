package matrix

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/tuple"
)

func TestZeroMatrix(t *testing.T) {
	m1 := ZeroMatrix()
	expected := Matrix{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	if m1 != expected {
		t.Errorf("expected %v, got %v", expected, m1)
	}
}

func TestIdentityMatrix(t *testing.T) {
	m1 := IdentityMatrix()
	expected := Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}

	if m1 != expected {
		t.Errorf("expected %v, got %v", expected, m1)
	}
}

func TestIdentityMatrixMultiplication(t *testing.T) {
	m1 := Matrix{
		{1, 2, 3, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	}

	result := m1.Multiply(IdentityMatrix())
	expected := Matrix{
		{1, 2, 3, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	}

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMatrixEquality(t *testing.T) {
	m1 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	m2 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	// Go is nice because arrays can be compared for value equality with != and ==
	// No need for a lame Equals method :)
	isEqual := m1 == m2
	expected := true

	if isEqual != expected {
		t.Errorf("expected %v, got %v", expected, isEqual)
	}
}

func TestMatrixMultiplication(t *testing.T) {
	m1 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	m2 := Matrix{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}
	m := m1.Multiply(m2)
	expected := Matrix{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}

	if m != expected {
		t.Errorf("expected %v, got %v", expected, m)
	}
}

func TestMatrixTupleMultiplication(t *testing.T) {
	m1 := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}

	operand := tuple.Tuple{X: 1, Y: 2, Z: 3, W: 1}
	result := m1.MultiplyTuple(operand)
	expected := tuple.Tuple{X: 18, Y: 24, Z: 33, W: 1}

	if !result.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMatrixTranspose(t *testing.T) {
	m1 := Matrix{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	}

	transposed := m1.Transpose()
	expected := Matrix{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	}

	if transposed != expected {
		t.Errorf("expected %v, got %v", expected, transposed)
	}
}

func TestTransposeIdentityMatrix(t *testing.T) {
	expected := IdentityMatrix()
	transposed := expected.Transpose()

	if transposed != expected {
		t.Errorf("expected %v, got %v", expected, transposed)
	}
}

func TestDeterminant(t *testing.T) {
	m := [2][2]float64{
		{1, 5},
		{-3, 2},
	}

	d := Determinant(m)
	var expected float64 = 17

	if d != expected {
		t.Errorf("expected %v, got %v", expected, d)
	}
}

func TestSubMatrix(t *testing.T) {
	m := Matrix{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	}

	sub := SubMatrix(m, 2, 1)
	expected := [][]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	}

	if !isEqual(sub, expected) {
		t.Errorf("expected %v, got %v", expected, sub)
	}

}

func isEqual(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
