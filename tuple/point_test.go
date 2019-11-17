package tuple

import "testing"

func TestNewPoint(t *testing.T) {
	tuple := Tuple{1, 2, 3, 1}
	point := NewPoint(1, 2, 3)

	if point.x != tuple.x ||
		point.y != tuple.y ||
		point.z != tuple.z ||
		point.w != tuple.w {
		t.Errorf("got %v, expected %v", point, tuple)
	}
}

func TestPointSubtractPoint(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)

	vector := p1.Subtract(p2)
	expected := NewVector(-2, -4, -6)

	if !vector.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, vector)
	}
}

func TestPointSubtractVector(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewVector(5, 6, 7)

	point := p1.Subtract(p2)
	expected := NewPoint(-2, -4, -6)

	if !point.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, point)
	}
}
