package math

import "testing"

func TestNewVector(t *testing.T) {
	tuple := Tuple{1, 2, 3, 0}
	point := Vector(1, 2, 3)

	if point.x != tuple.x ||
		point.y != tuple.y ||
		point.z != tuple.z ||
		point.w != tuple.w {
		t.Errorf("got %v, expected %v", point, tuple)
	}
}

func TestVectorSubtractVector(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)

	vector := v1.subtract(v2)
	expected := Vector(-2, -4, -6)

	if !vector.equals(expected) {
		t.Errorf("expected %v, got %v", expected, vector)
	}
}
