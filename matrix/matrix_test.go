package matrix

import (
	"math"
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

	if !m1.Equals(expected) {
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

	if !m1.Equals(expected) {
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

	if !result.Equals(expected) {
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

	isEqual := m1.Equals(m2)
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

	if !m.Equals(expected) {
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

	if !transposed.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, transposed)
	}
}

func TestTransposeIdentityMatrix(t *testing.T) {
	expected := IdentityMatrix()
	transposed := expected.Transpose()

	if !transposed.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, transposed)
	}
}

func TestDeterminant_1(t *testing.T) {
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

func TestDeterminant_2(t *testing.T) {
	m := Matrix{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}

	d := m.Determinant()
	var expected float64 = -196

	if d != expected {
		t.Errorf("expected %v, got %v", expected, d)
	}
}

func TestDeterminant_3(t *testing.T) {
	m := Matrix{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}

	d := m.Determinant()
	var expected float64 = -4071

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

	if !sub.Equals(expected) {
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
		t.Errorf("expected %v, got %v", expected, cofactor)
	}
}

func TestCofactor_3(t *testing.T) {
	m := Matrix{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}
	cofactor := m.Cofactor(0, 2)
	var expected float64 = -46

	if cofactor != expected {
		t.Errorf("expected %v, got %v", expected, cofactor)
	}
}

func TestInvertible_1(t *testing.T) {
	m := Matrix{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	}

	isInvertible := m.IsInvertible()
	expected := true

	if isInvertible != expected {
		t.Errorf("expected %v, got %v", expected, isInvertible)
	}
}

func TestInvertible_2(t *testing.T) {
	m := Matrix{
		{-4, 2, -2, 3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	}

	isInvertible := m.IsInvertible()
	expected := false

	if isInvertible != expected {
		t.Errorf("expected %v, got %v", expected, isInvertible)
	}
}

func TestInvert_1(t *testing.T) {
	m := Matrix{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	}

	inverse := m.Invert()

	d := m.Determinant()
	cof2 := m.Cofactor(3, 2)

	r := inverse[2][3]
	expected := cof2 / d

	if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestInvert_2(t *testing.T) {
	m := Matrix{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	}

	inverse := m.Invert()
	expected := Matrix{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	}

	if !inverse.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, inverse)
	}
}

func TestInvert_3(t *testing.T) {
	m := Matrix{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	}

	inverse := m.Invert()
	expected := Matrix{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	}

	if !inverse.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, inverse)
	}
}

func TestInvert_4(t *testing.T) {
	m := Matrix{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	}

	inverse := m.Invert()
	expected := Matrix{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	}

	if !inverse.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, inverse)
	}
}

func TestMultiplyProductByInverse(t *testing.T) {
	m1 := Matrix{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	}

	m2 := Matrix{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	}

	product := m1.Multiply(m2)
	original := product.Multiply(m2.Invert())
	isEqual := m1.Equals(original)
	expected := true

	if isEqual != true {
		t.Errorf("expected %v, got %v", expected, isEqual)
	}
}

func TestTranslatePoint(t *testing.T) {
	translation := TranslationMatrix(5, -3, 2)
	p := tuple.NewPoint(-3, 4, 5)

	transform := translation.MultiplyTuple(p)
	expected := tuple.NewPoint(2, 1, 7)

	if !transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, transform)
	}
}

func TestInverseTranslatePoint(t *testing.T) {
	translation := TranslationMatrix(5, -3, 2)
	inv := translation.Invert()
	p := tuple.NewPoint(-3, 4, 5)

	transform := inv.MultiplyTuple(p)
	expected := tuple.NewPoint(-8, 7, 3)

	if !transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, transform)
	}
}

func TestTranslateVector(t *testing.T) {
	translation := TranslationMatrix(5, -3, 2)
	inv := translation.Invert()
	v := tuple.NewVector(-3, 4, 5)

	transform := inv.MultiplyTuple(v)
	expected := tuple.NewVector(-3, 4, 5)

	if !transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, transform)
	}
}

func TestScalingPoint(t *testing.T) {
	scalar := ScalingMatrix(2, 3, 4)
	p := tuple.NewPoint(-4, 6, 8)

	scaled := scalar.MultiplyTuple(p)
	expected := tuple.NewPoint(-8, 18, 32)

	if !scaled.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, scaled)
	}
}

func TestScalingVector(t *testing.T) {
	scalar := ScalingMatrix(2, 3, 4)
	v := tuple.NewVector(-4, 6, 8)

	scaled := scalar.MultiplyTuple(v)
	expected := tuple.NewVector(-8, 18, 32)

	if !scaled.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, scaled)
	}
}

func TestInverseScalingVector(t *testing.T) {
	scalar := ScalingMatrix(2, 3, 4)
	v := tuple.NewVector(-4, 6, 8)

	scaled := scalar.Invert().MultiplyTuple(v)
	expected := tuple.NewVector(-2, 2, 2)

	if !scaled.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, scaled)
	}
}

// AKA Reflection
func TestNegativeScalingPoint(t *testing.T) {
	scalar := ScalingMatrix(-1, 1, 1)
	p := tuple.NewPoint(2, 3, 4)

	scaled := scalar.MultiplyTuple(p)
	expected := tuple.NewPoint(-2, 3, 4)

	if !scaled.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, scaled)
	}
}

func TestXAxisRotation_1(t *testing.T) {
	p := tuple.NewPoint(0, 1, 0)
	axis := tuple.Right()
	halfQuarter := RotationMatrix(axis, math.Pi/4)

	rotated := halfQuarter.MultiplyTuple(p)
	expected := tuple.NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	if !rotated.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, rotated)
	}
}

func TestXAxisRotation_2(t *testing.T) {
	p := tuple.NewPoint(0, 1, 0)
	axis := tuple.Right()
	fullQuarter := RotationMatrix(axis, math.Pi/2)

	rotated := fullQuarter.MultiplyTuple(p)
	expected := tuple.NewPoint(0, 0, 1)

	if !rotated.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, rotated)
	}
}

func TestInverseXRotation(t *testing.T) {
	p := tuple.NewPoint(0, 1, 0)
	axis := tuple.Right()
	halfQuarter := RotationMatrix(axis, math.Pi/4)

	inverse := halfQuarter.Invert().MultiplyTuple(p)
	expected := tuple.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)

	if !inverse.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, inverse)
	}
}
func TestYAxisRotation_1(t *testing.T) {
	p := tuple.NewPoint(0, 0, 1)
	axis := tuple.Up()
	halfQuarter := RotationMatrix(axis, math.Pi/4)

	rotated := halfQuarter.MultiplyTuple(p)
	expected := tuple.NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)

	if !rotated.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, rotated)
	}
}

func TestYAxisRotation_2(t *testing.T) {
	p := tuple.NewPoint(0, 0, 1)
	axis := tuple.Up()
	fullQuarter := RotationMatrix(axis, math.Pi/2)

	rotated := fullQuarter.MultiplyTuple(p)
	expected := tuple.NewPoint(1, 0, 0)

	if !rotated.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, rotated)
	}
}

func TestZAxisRotation_1(t *testing.T) {
	p := tuple.NewPoint(0, 1, 0)
	axis := tuple.Back()
	halfQuarter := RotationMatrix(axis, math.Pi/4)

	rotated := halfQuarter.MultiplyTuple(p)
	expected := tuple.NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

	if !rotated.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, rotated)
	}
}

func TestZAxisRotation_2(t *testing.T) {
	p := tuple.NewPoint(0, 1, 0)
	axis := tuple.Back()
	fullQuarter := RotationMatrix(axis, math.Pi/2)

	rotated := fullQuarter.MultiplyTuple(p)
	expected := tuple.NewPoint(-1, 0, 0)

	if !rotated.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, rotated)
	}
}
