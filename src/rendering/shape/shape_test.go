package shape

import (
	"math"
	"testing"

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
