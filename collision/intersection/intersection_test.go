package intersection

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/calculation/point"
	"github.com/kingsleyliao/ray-tracer/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/collision/ray"
	"github.com/kingsleyliao/ray-tracer/rendering/shape"
)

// An Intersection encapsulates t and object
func TestIntersection_1(t *testing.T) {
	s := shape.NewSphere()
	i := NewIntersection(3.5, s)

	expectedT := 3.5
	expectedS := s

	if expectedT != i.T || !i.Object.Equals(expectedS) {
		t.Errorf("expected %v, %v, got %v, %v", expectedT, expectedS, i.T, i.Object)
	}
}

// A Ray intersects a Sphere at two Points
func TestIntersect_1(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()

	xs := Intersect(s, r)

	expected0 := s.Equals(xs[0].Object)
	expected1 := s.Equals(xs[1].Object)

	if expected0 != true || expected1 != true {
		t.Errorf("expected %v, %v, got %v, %v", expected0, expected1, xs[0], xs[1])
	}
}

// A Ray intersects a Sphere at a tangent
func TestIntersect_2(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 1, -5), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()

	xs := Intersect(s, r)

	expected0 := s.Equals(xs[0].Object)
	expected1 := s.Equals(xs[1].Object)
	expected2 := len(xs) == 2

	if expected0 != true || expected1 != true || expected2 != true {
		t.Errorf("expected %v, %v, %v got %v, %v, %v", expected0, expected1, expected2, xs[0], xs[1], len(xs))
	}
}

// A Ray misses a Sphere
func TestIntersect_3(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 2, -5), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()

	xs := Intersect(s, r)

	l := len(xs)
	expected := 0

	if l != expected {
		t.Errorf("expected %v, got %v", expected, l)
	}
}

// A Ray originates in a Sphere
func TestIntersect_4(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()

	xs := Intersect(s, r)

	expected := len(xs) == 2
	expected1 := s.Equals(xs[0].Object)
	expected2 := s.Equals(xs[1].Object)

	if expected != true || expected1 != true || expected2 != true {
		t.Errorf("expected %v, %v, %v got %v, %v, %v", expected, expected1, expected2, len(xs), xs[0], xs[1])
	}
}

// A Sphere is behind a Ray
func TestIntersect_5(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, 5), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()

	xs := Intersect(s, r)

	expected := len(xs) == 2
	expected1 := s.Equals(xs[0].Object)
	expected2 := s.Equals(xs[1].Object)

	if expected != true || expected1 != true || expected2 != true {
		t.Errorf("expected %v, %v, %v got %v, %v, %v", expected, expected1, expected2, len(xs), xs[0], xs[1])
	}
}

// The Hit, when all Intersections have positive T
func TestHit_1(t *testing.T) {
	s := shape.NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := Intersections(i1, i2)

	i, _ := Hit(xs)

	if !i.Equals(i1) {
		t.Errorf("expected %v, got %v", i1, i)
	}
}

// The Hit, when some Intersections have negative T
func TestHit_2(t *testing.T) {
	s := shape.NewSphere()
	i1 := NewIntersection(-1, s)
	i2 := NewIntersection(1, s)
	xs := Intersections(i2, i1)

	i, _ := Hit(xs)

	if !i.Equals(i2) {
		t.Errorf("expected %v, got %v", i2, i)
	}
}

// The Hit, when all Intersections have negative T
func TestHit_3(t *testing.T) {
	s := shape.NewSphere()
	i1 := NewIntersection(-2, s)
	i2 := NewIntersection(-1, s)
	xs := Intersections(i2, i1)

	_, ok := Hit(xs)

	if ok != false {
		t.Errorf("expected %v, got %v", false, ok)
	}
}

// The Hit is always the lowest non-negative intersection
func TestHit_4(t *testing.T) {
	s := shape.NewSphere()
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs := Intersections(i1, i2, i3, i4)

	i, _ := Hit(xs)

	if !i.Equals(i4) {
		t.Errorf("expected %v, got %v", i4, i)
	}
}

// Intersecting a scaled Sphere with a ray
func TestIntersectingSphere_1(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()
	s.SetTransform(matrix.ScalingMatrix(2, 2, 2))
	xs := Intersect(s, r)

	expectedCount := 2
	expectedT1 := 3.0
	expectedT2 := 7.0

	if expectedCount != len(xs) {
		t.Errorf("expected %v, got %v", expectedCount, len(xs))
	}

	if expectedT1 != xs[0].T {
		t.Errorf("expected %v, got %v", expectedT1, xs[0].T)
	}

	if expectedT2 != xs[1].T {
		t.Errorf("expected %v, got %v", expectedT2, xs[1].T)
	}
}

// Intersecting a translated Sphere with a ray
func TestIntersectingSphere_2(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()
	s.SetTransform(matrix.TranslationMatrix(5, 0, 0))
	xs := Intersect(s, r)

	expected := 0

	if expected != len(xs) {
		t.Errorf("expected %v, got %v", expected, len(xs))
	}
}
