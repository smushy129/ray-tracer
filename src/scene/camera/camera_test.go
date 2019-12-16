package camera

import (
	"math"
	"testing"
)

func TestNewCamera(t *testing.T) {
	hsize := 160
	vsize := 120
	fov := math.Pi / 2

	c := NewCamera(hsize, vsize, fov)

	if c.Hsize != 160 {
		t.Errorf("expected %v, got %v", 160, c.Hsize)
	}

	if c.Vsize != 120 {
		t.Errorf("expected %v, got %v", 120, c.Vsize)
	}

	if c.Fov != math.Pi/2 {
		t.Errorf("expected %v, got %v", math.Pi/2, c.Fov)
	}
}
