package vector

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/src/calculation/tuple"
)

func TestNewVector(t *testing.T) {
	tuple := tuple.Tuple{X: 1, Y: 2, Z: 3, W: 0}
	vector := NewVector(1, 2, 3)

	if vector.X != tuple.X ||
		vector.Y != tuple.Y ||
		vector.Z != tuple.Z ||
		vector.W != tuple.W {
		t.Errorf("got %v, expected %v", vector, tuple)
	}
}
