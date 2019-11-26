package tuple

import (
	"math"
	"testing"
)

func TestTupleEquals__true1(t *testing.T) {
	t1 := Tuple{1, 2, 3, 1}
	t2 := Tuple{1, 2, 3, 1}

	isEqual := t1.Equals(t2)

	if isEqual != true {
		t.Errorf("got %v, expected %v", isEqual, true)
	}
}

func TestTupleEquals__true2(t *testing.T) {
	t1 := Tuple{1, 2, 3.000000001, 1}
	t2 := Tuple{1, 2, 3, 1}

	isEqual := t1.Equals(t2)

	if isEqual != true {
		t.Errorf("got %v, expected %v", isEqual, true)
	}
}

func TestTupleEquals__false(t *testing.T) {
	t1 := Tuple{1, 2, 3, 0}
	t2 := Tuple{1, 2, 3, 1}

	isEqual := t1.Equals(t2)

	if isEqual != false {
		t.Errorf("got %v, expected %v", isEqual, false)
	}
}

func TestTupleAdd(t *testing.T) {
	t1 := Tuple{3, -2, 5, 1}
	t2 := Tuple{-2, 3, 1, 0}

	sum := t1.Add(t2)
	expected := Tuple{1, 1, 6, 1}

	if !sum.Equals(expected) {
		t.Errorf("got %v, expected %v", sum, expected)
	}
}

func TestTupleInvert(t *testing.T) {
	t1 := Tuple{1, 2, 3, -4}

	inverted := t1.Invert()
	expected := Tuple{-1, -2, -3, 4}

	if !inverted.Equals(expected) {
		t.Errorf("got %v, expected %v", inverted, expected)
	}
}

func TestTupleScale(t *testing.T) {
	t1 := Tuple{1, 2, 3, -4}

	scaled := t1.Scale(2)
	expected := Tuple{2, 4, 6, -8}

	if !scaled.Equals(expected) {
		t.Errorf("got %v, expected %v", scaled, expected)
	}
}

func TestVectorSubtractVector(t *testing.T) {
	v1 := Tuple{3, 2, 1, 0}
	v2 := Tuple{5, 6, 7, 0}

	vector := v1.Subtract(v2)
	// Vector subtracted from another vector returns a vector, hence W:0
	expected := Tuple{-2, -4, -6, 0}

	if !vector.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, vector)
	}
}

func TestTupleMagnitude_1(t *testing.T) {
	v := Tuple{1, 0, 0, 0}
	var expected float64 = 1

	if v.Magnitude() != expected {
		t.Errorf("got %v, expected %v", v.Magnitude(), expected)
	}
}

func TestVectorMagnitude_2(t *testing.T) {
	v := Tuple{0, 1, 0, 0}
	var expected float64 = 1

	if v.Magnitude() != expected {
		t.Errorf("got %v, expected %v", v.Magnitude(), expected)
	}
}

func TestVectorMagnitude_3(t *testing.T) {
	v := Tuple{0, 0, 1, 0}
	var expected float64 = 1

	if v.Magnitude() != expected {
		t.Errorf("got %v, expected %v", v.Magnitude(), expected)
	}
}

func TestVectorMagnitude_4(t *testing.T) {
	v := Tuple{1, 2, 3, 0}
	var expected = math.Sqrt(14)

	if v.Magnitude() != expected {
		t.Errorf("got %v, expected %v", v.Magnitude(), expected)
	}
}

func TestVectorMagnitude_5(t *testing.T) {
	v := Tuple{-1, -2, -3, 0}
	var expected = math.Sqrt(14)

	if v.Magnitude() != expected {
		t.Errorf("got %v, expected %v", v.Magnitude(), expected)
	}
}

// Magnitude of a normalized vector is 1
func TestVectorMagnitude_6(t *testing.T) {
	v := Tuple{1, 2, 3, 0}
	norm := v.Normalize()
	var expected float64 = 1

	if norm.Magnitude() != expected {
		t.Errorf("got %v, expected %v", norm, expected)
	}
}

// Normalizing a vector returns a unit vector
func TestVectorNormalize_1(t *testing.T) {
	v := Tuple{4, 0, 0, 0}

	expected := Tuple{1, 0, 0, 0}

	if v.Normalize() != expected {
		t.Errorf("got %v, expected %v", v.Normalize(), expected)
	}
}

// Normalizing a vector returns a unit vector
func TestVectorNormalize_2(t *testing.T) {
	v := Tuple{1, 2, 3, 0}
	m := math.Sqrt(14)
	expected := Tuple{1 / m, 2 / m, 3 / m, 0}

	if v.Normalize() != expected {
		t.Errorf("got %v, expected %v", v.Normalize(), expected)
	}

}

func TestDotProduct(t *testing.T) {
	v1 := Tuple{1, 2, 3, 0}
	v2 := Tuple{2, 3, 4, 0}

	dp := v1.Dot(v2)
	// Dot product of two Vectors is a float
	var expected float64 = 20

	if dp != expected {
		t.Errorf("got %v, expected %v", dp, expected)
	}
}

func TestCrossProduct_1(t *testing.T) {
	v1 := Tuple{1, 2, 3, 0}
	v2 := Tuple{2, 3, 4, 0}

	c := v1.Cross(v2)
	// Cross product of two Vectors is another Vector
	expected := Tuple{-1, 2, -1, 0}

	if c != expected {
		t.Errorf("got %v, expected %v", c, expected)
	}
}

func TestCrossProduct_2(t *testing.T) {
	v1 := Tuple{1, 2, 3, 0}
	v2 := Tuple{2, 3, 4, 0}

	c := v2.Cross(v1)
	// Cross product of two Vectors is another Vector
	expected := Tuple{1, -2, 1, 0}

	if c != expected {
		t.Errorf("got %v, expected %v", c, expected)
	}
}
