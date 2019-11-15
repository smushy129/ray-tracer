package math

import "testing"

func TestNewPoint(t *testing.T) {
	tuple := Tuple{1, 2, 3, 1}
	point := Point(1, 2, 3)

	if point.x != tuple.x ||
		point.y != tuple.y ||
		point.z != tuple.z ||
		point.w != tuple.w {
		t.Errorf("got %v, expected %v", point, tuple)
	}
}

func TestPointSubtractPoint(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)

	vector := p1.subtract(p2)
	expected := Vector(-2, -4, -6)

	if !vector.equals(expected) {
		t.Errorf("expected %v, got %v", expected, vector)
	}
}

func TestPointSubtractVector(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Vector(5, 6, 7)

	point := p1.subtract(p2)
	expected := Point(-2, -4, -6)

	if !point.equals(expected) {
		t.Errorf("expected %v, got %v", expected, point)
	}
}
