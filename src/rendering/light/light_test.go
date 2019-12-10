package light

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/src/rendering/material"

	"github.com/kingsleyliao/ray-tracer/src/calculation/point"

	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
)

// A point light has a position and intensity
func TestNewPointLight(t *testing.T) {
	i := color.NewColor(1, 1, 1)
	p := point.Zero()

	l := NewPointLight(p, i)

	expectedPoint := l.Position == p
	expectedIntensity := l.Intensity == i

	if expectedPoint != true || expectedIntensity != true {
		t.Errorf("expected %v, %v, got %v, %v", expectedPoint, expectedIntensity, l.Position, l.Intensity)
	}
}

// Lighting with the eye between the light and the surface
func TestLighting_1(t *testing.T) {
	m := material.NewMaterial()
	p := point.Zero()
	eyeV := vector.NewVector(0, 0, -1)
	normalV := vector.NewVector(0, 0, -1)

	light := NewPointLight(point.NewPoint(0, 0, -10), color.NewColor(1, 1, 1))
	result := Lighting(m, light, p, eyeV, normalV)

	expected := color.NewColor(1.9, 1.9, 1.9)

	if expected != result {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
