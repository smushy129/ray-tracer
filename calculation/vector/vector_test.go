package vector

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/calculation/tuple"
)

func TestNewVector(t *testing.T) {
	tuple := tuple.Tuple{X: 1, Y: 2, Z: 3, W: 0}
	point := NewVector(1, 2, 3)

	if point.X != tuple.X ||
		point.Y != tuple.Y ||
		point.Z != tuple.Z ||
		point.W != tuple.W {
		t.Errorf("got %v, expected %v", point, tuple)
	}
}
