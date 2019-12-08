package point

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/src/calculation/tuple"
	vector "github.com/kingsleyliao/ray-tracer/src/calculation/vector"
)

func TestNewPoint(t *testing.T) {
	point := NewPoint(1, 2, 3)
	expected := tuple.Tuple{X: 1, Y: 2, Z: 3, W: 1}

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
	// Point subtracted from a point gives a vector, hence W: 0
	expected := tuple.Tuple{X: -2, Y: -4, Z: -6, W: 0}

	if !vector.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, vector)
	}
}

func TestPointSubtractVector(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := vector.NewVector(5, 6, 7)

	point := p1.Subtract(p2)
	// Vector subtracted from a point gives a new point, hence W: 1
	expected := NewPoint(-2, -4, -6)

	if !point.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, point)
	}
}
