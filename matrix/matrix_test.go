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
