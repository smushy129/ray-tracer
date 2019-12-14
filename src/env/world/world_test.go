package world

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
	"github.com/kingsleyliao/ray-tracer/src/rendering/light"
	"github.com/kingsleyliao/ray-tracer/src/rendering/material"
	"github.com/kingsleyliao/ray-tracer/src/rendering/shape"
)

func TestDefaultWorld(t *testing.T) {
	w := DefaultWorld()

	s1 := shape.NewSphere()
	s1.Material = material.Material{
		Color:    color.NewColor(0.8, 1.0, 0.6),
		Diffuse:  0.7,
		Specular: 0.2,
	}
	s2 := shape.NewSphere()
	s2.Transform = matrix.ScalingMatrix(0.5, 0.5, 0.5)

	l := light.NewPointLight(point.NewPoint(-10, 10, -10), color.NewColor(1, 1, 1))

	if !l.Equals(w.Light) {
		t.Errorf("expected %v, got %v", true, false)
	}

	if !s1.Equals(w.Objects[0]) {
		t.Errorf("expected %v, got %v", true, false)
	}

	if !s2.Equals(w.Objects[1]) {
		t.Errorf("expected %v, got %v", true, false)
	}
}

// Intersecting a world with a ray
func TestIntersectRayWithWorld(t *testing.T) {

}
