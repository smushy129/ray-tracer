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

	if !m1.IsEqual(expected) {
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

	if !m1.IsEqual(expected) {
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

	if !result.IsEqual(expected) {
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

	IsEqual := m1.IsEqual(m2)
	expected := true

	if IsEqual != expected {
		t.Errorf("expected %v, got %v", expected, IsEqual)
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

	if !m.IsEqual(expected) {
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

	if !transposed.IsEqual(expected) {
		t.Errorf("expected %v, got %v", expected, transposed)
	}
}

func TestTransposeIdentityMatrix(t *testing.T) {
	expected := IdentityMatrix()
	transposed := expected.Transpose()

	if !transposed.IsEqual(expected) {
		t.Errorf("expected %v, got %v", expected, transposed)
	}
}

func TestDeterminant(t *testing.T) {
	m := Matrix{
		{1, 5},
		{-3, 2},
	}

	d := m.Determinant()
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

	sub := m.SubMatrix(2, 1)
	expected := [][]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	}

	if !sub.IsEqual(expected) {
		t.Errorf("expected %v, got %v", expected, sub)
	}

}

func TestMinor_1(t *testing.T) {
	m := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	sub := m.SubMatrix(1, 0)
	d := sub.Determinant()
	minor := m.Minor(1, 0)

	if d != minor {
		t.Errorf("expected %v, got %v", d, minor)
	}
}

func TestMinor_2(t *testing.T) {
	m := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}
	minor := m.Minor(0, 0)
	var expected float64 = -12

	if minor != expected {
		t.Errorf("expected %v, got %v", expected, minor)
	}
}

func TestMinor_3(t *testing.T) {
	m := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}
	minor := m.Minor(1, 0)
	var expected float64 = 25

	if minor != expected {
		t.Errorf("expected %v, got %v", expected, minor)
	}
}

func TestCofactor_1(t *testing.T) {
	m := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}
	cofactor := m.Cofactor(0, 0)
	var expected float64 = -12

	if cofactor != expected {
		t.Errorf("expected %v, got %v", expected, cofactor)
	}
}

func TestCofactor_2(t *testing.T) {
	m := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}
	cofactor := m.Cofactor(1, 0)
	var expected float64 = -25

	if cofactor != expected {
		cofactor := m.Cofactor(1, 0)
		t.Errorf("expected %v, got %v", expected, cofactor)
	}
}
