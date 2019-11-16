package math

import (
	"math"
	"testing"
)

func TestNewVector(t *testing.T) {
	tuple := Tuple{1, 2, 3, 0}
	point := NewVector(1, 2, 3)

	if point.x != tuple.x ||
		point.y != tuple.y ||
		point.z != tuple.z ||
		point.w != tuple.w {
		t.Errorf("got %v, expected %v", point, tuple)
	}
}

func TestVectorSubtractVector(t *testing.T) {
	v1 := NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)

	vector := v1.subtract(v2)
	expected := NewVector(-2, -4, -6)

	if !vector.equals(expected) {
		t.Errorf("expected %v, got %v", expected, vector)
	}
}

func TestTupleMagnitude_1(t *testing.T) {
	v := NewVector(1, 0, 0)
	var expected float64 = 1

	if v.magnitude() != expected {
		t.Errorf("got %v, expected %v", v.magnitude(), expected)
	}
}

func TestVectorMagnitude_2(t *testing.T) {
	v := NewVector(0, 1, 0)
	var expected float64 = 1

	if v.magnitude() != expected {
		t.Errorf("got %v, expected %v", v.magnitude(), expected)
	}
}

func TestVectorMagnitude_3(t *testing.T) {
	v := NewVector(0, 0, 1)
	var expected float64 = 1

	if v.magnitude() != expected {
		t.Errorf("got %v, expected %v", v.magnitude(), expected)
	}
}

func TestVectorMagnitude_4(t *testing.T) {
	v := NewVector(1, 2, 3)
	var expected = math.Sqrt(14)

	if v.magnitude() != expected {
		t.Errorf("got %v, expected %v", v.magnitude(), expected)
	}
}

func TestVectorMagnitude_5(t *testing.T) {
	v := NewVector(-1, -2, -3)
	var expected = math.Sqrt(14)

	if v.magnitude() != expected {
		t.Errorf("got %v, expected %v", v.magnitude(), expected)
	}
}

func TestVectorMagnitude_6(t *testing.T) {
	v := NewVector(1, 2, 3)
	norm := v.normalize()
	var expected float64 = 1

	if norm.magnitude() != expected {
		t.Errorf("got %v, expected %v", norm, expected)
	}
}

func TestVectorNormalize_1(t *testing.T) {
	v := NewVector(4, 0, 0)
	expected := NewVector(1, 0, 0)

	if v.normalize() != expected {
		t.Errorf("got %v, expected %v", v.normalize(), expected)
	}
}

func TestVectorNormalize_2(t *testing.T) {
	v := NewVector(1, 2, 3)
	m := math.Sqrt(14)
	expected := NewVector(1/m, 2/m, 3/m)

	if v.normalize() != expected {
		t.Errorf("got %v, expected %v", v.normalize(), expected)
	}

}

func TestDotProduct(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	dp := v1.dot(v2)
	var expected float64 = 20

	if dp != expected {
		t.Errorf("got %v, expected %v", dp, expected)
	}
}

func TestCrossProduct_1(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	c := v1.cross(v2)
	expected := NewVector(-1, 2, -1)

	if c != expected {
		t.Errorf("got %v, expected %v", c, expected)
	}
}

func TestCrossProduct_2(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	c := v2.cross(v1)
	expected := NewVector(1, -2, 1)

	if c != expected {
		t.Errorf("got %v, expected %v", c, expected)
	}
}
