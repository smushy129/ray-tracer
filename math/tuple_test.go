package math

import "testing"

func TestPoint(t *testing.T) {
	tuple := Tuple{1, 2, 3, 1}
	point := Point(1, 2, 3)

	if point.x != tuple.x ||
		point.y != tuple.y ||
		point.z != tuple.z ||
		point.w != tuple.w {
		t.Errorf("got %v, expected %v", point, tuple)
	}
}

func TestVector(t *testing.T) {
	tuple := Tuple{1, 2, 3, 0}
	point := Vector(1, 2, 3)

	if point.x != tuple.x ||
		point.y != tuple.y ||
		point.z != tuple.z ||
		point.w != tuple.w {
		t.Errorf("got %v, expected %v", point, tuple)
	}
}
