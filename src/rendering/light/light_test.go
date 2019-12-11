package light

import (
	"math"
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

	l := NewPointLight(point.NewPoint(0, 0, -10), color.NewColor(1, 1, 1))
	result := Lighting(m, l, p, eyeV, normalV)

	expected := color.NewColor(1.9, 1.9, 1.9)

	if !expected.Equals(result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Lighting with the eye between the light and surface, eye offset 45˚
func TestLighting_2(t *testing.T) {
	m := material.NewMaterial()
	p := point.Zero()
	eyeV := vector.NewVector(0, math.Sqrt(2)/2, math.Sqrt(2)/2)
	normalV := vector.NewVector(0, 0, -1)

	l := NewPointLight(point.NewPoint(0, 0, -10), color.NewColor(1, 1, 1))
	result := Lighting(m, l, p, eyeV, normalV)

	expected := color.NewColor(1, 1, 1)

	if !expected.Equals(result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Lighting with eye opposite surface, light offset 45˚
func TestLighting_3(t *testing.T) {
	m := material.NewMaterial()
	p := point.Zero()
	eyeV := vector.NewVector(0, 0, -1)
	normalV := vector.NewVector(0, 0, -1)

	l := NewPointLight(point.NewPoint(0, 10, -10), color.NewColor(1, 1, 1))
	result := Lighting(m, l, p, eyeV, normalV)

	expected := color.NewColor(0.7364, 0.7364, 0.7364)

	if !expected.Equals(result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Lighting with eye in the path of the reflection vector
func TestLighting_4(t *testing.T) {
	m := material.NewMaterial()
	p := point.Zero()
	eyeV := vector.NewVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalV := vector.NewVector(0, 0, -1)

	l := NewPointLight(point.NewPoint(0, 10, -10), color.NewColor(1, 1, 1))
	result := Lighting(m, l, p, eyeV, normalV)

	expected := color.NewColor(1.6364, 1.6364, 1.6364)

	if !expected.Equals(result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Lighting with the light behind the surface
func TestLighting_5(t *testing.T) {
	m := material.NewMaterial()
	p := point.Zero()
	eyeV := vector.NewVector(0, 0, -1)
	normalV := vector.NewVector(0, 0, -1)

	l := NewPointLight(point.NewPoint(0, 0, 10), color.NewColor(1, 1, 1))
	result := Lighting(m, l, p, eyeV, normalV)

	expected := color.NewColor(0.1, 0.1, 0.1)

	if !expected.Equals(result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
