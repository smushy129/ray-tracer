package tuple

import "testing"

func TestNewPoint(t *testing.T) {
	point := NewPoint(1, 2, 3)
	expected := Tuple{1, 2, 3, 1}

	if point.X != expected.X ||
		point.Y != expected.Y ||
		point.Z != expected.Z ||
		point.W != expected.W {
		t.Errorf("expected %v, got %v", expected, point)
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
