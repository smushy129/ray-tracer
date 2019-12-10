package material

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
)

func TestNewMaterial(t *testing.T) {
	m := NewMaterial()

	if m.Color != color.NewColor(1, 1, 1) {
		t.Errorf("expected %v, got %v", color.NewColor(1, 1, 1), m.Color)
	}

	if m.Ambient != 0.1 {
		t.Errorf("expected %v, got %v", 0.1, m.Ambient)
	}

	if m.Diffuse != 0.9 {
		t.Errorf("expected %v, got %v", 0.9, m.Diffuse)
	}

	if m.Specular != 0.9 {
		t.Errorf("expected %v, got %v", 0.9, m.Specular)
	}

	if m.Shininess != 200.0 {
		t.Errorf("expected %v, got %v", 0.9, m.Shininess)
	}
}
