package shape

import (
	"math"
	"testing"

	"github.com/kingsleyliao/ray-tracer/src/rendering/material"

	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"

	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
)

// Creating a new Sphere has a default Transform of the IdentityMatrix
func TestNewSphere(t *testing.T) {
	s := NewSphere()

	expected := matrix.IdentityMatrix()

	if !s.Transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, s.Transform)
	}
}

// Creating a sphere has a default Material
func TestDefaultMaterial(t *testing.T) {
	s := NewSphere()
	m := s.Material

	expected := material.DefaultMaterial()

	if m != expected {
		t.Errorf("expected %v, got %v", expected, m)
	}
}

// A Sphere can be assigned a Material
func TestAssignMaterial(t *testing.T) {
	s := NewSphere()
	m := material.DefaultMaterial()
	m.Ambient = 1
	s.Material = m

	expected := m

	if expected != s.Material {
		t.Errorf("expected %v, got %v", expected, expected == s.Material)
	}
}

func TestTransformSphere(t *testing.T) {
	s := NewSphere()
	m := matrix.TranslationMatrix(2, 3, 4)
	s.SetTransform(m)

	expected := matrix.TranslationMatrix(2, 3, 4)

	if !s.Transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, s.Transform)
	}
}

// The Normal on a Sphere at a point on the x-axis
func TestNormalAt_1(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(point.NewPoint(1, 0, 0))

	expected := vector.NewVector(1, 0, 0)

	if !n.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, n)
	}
}

// The Normal on a Sphere at a point on the y-axis
func TestNormalAt_2(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(point.NewPoint(0, 1, 0))

	expected := vector.NewVector(0, 1, 0)

	if !n.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, n)
	}
}

// The Normal on a Sphere at a point on the z-axis
func TestNormalAt_3(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(point.NewPoint(0, 0, 1))

	expected := vector.NewVector(0, 0, 1)

	if !n.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, n)
	}
}

// The Normal on a Sphere at a point at a non-axial point
func TestNormalAt_4(t *testing.T) {
	s := NewSphere()
	v := math.Sqrt(3) / 3
	n := s.NormalAt(point.NewPoint(v, v, v))

	expected := vector.NewVector(v, v, v)

	if !n.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, n)
	}
}

// Computing the normal on a translated Sphere
func TestNormalAt_5(t *testing.T) {
	s := NewSphere()
	s.SetTransform(matrix.TranslationMatrix(0, 1, 0))
	n := s.NormalAt(point.NewPoint(0, 1.70711, -0.70711))

	expected := vector.NewVector(0, 0.70711, -0.70711)

	if !n.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, n)
	}
}

// Computing the normal on a transformed Sphere
func TestNormalAt_6(t *testing.T) {
	s := NewSphere()
	m := matrix.ScalingMatrix(1, 0.5, 1).Multiply(matrix.RotationMatrix(vector.Back(), math.Pi/5))
	s.SetTransform(m)

	v := math.Sqrt(2) / 2
	n := s.NormalAt(point.NewPoint(0, v, -v))

	expected := vector.NewVector(0, 0.97014, -0.24254)

	if !n.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, n)
	}
}
