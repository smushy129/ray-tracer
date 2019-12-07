package intersection

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/calculation/point"
	"github.com/kingsleyliao/ray-tracer/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/collision/ray"
	"github.com/kingsleyliao/ray-tracer/rendering/shape"
)

// A ray intersects a sphere at two points
func TestIntersect_1(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()

	xs := Intersect(s, r)

	expected0 := 4.0 == xs[0]
	expected1 := 6.0 == xs[1]

	if expected0 != true || expected1 != true {
		t.Errorf("expected %v, %v, got %v, %v", expected0, expected1, xs[0], xs[1])
	}
}

// A ray intersects a sphere at a tangent
func TestIntersect_2(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 1, -5), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()

	xs := Intersect(s, r)

	expected0 := 5.0 == xs[0]
	expected1 := 5.0 == xs[1]
	expected2 := len(xs) == 2

	if expected0 != true || expected1 != true || expected2 != true {
		t.Errorf("expected %v, %v, %v got %v, %v, %v", expected0, expected1, expected2, xs[0], xs[1], len(xs))
	}
}

// A ray misses a sphere
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

// A ray originates in a sphere
func TestIntersect_4(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()

	xs := Intersect(s, r)

	expected := len(xs) == 2
	expected1 := -1.0 == xs[0]
	expected2 := 1.0 == xs[1]

	if expected != true || expected1 != true || expected2 != true {
		t.Errorf("expected %v, %v, %v got %v, %v, %v", expected, expected1, expected2, len(xs), xs[0], xs[1])
	}
}

// A sphere is behind a ray
func TestIntersect_5(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, 5), vector.NewVector(0, 0, 1))
	s := shape.NewSphere()

	xs := Intersect(s, r)

	expected := len(xs) == 2
	expected1 := -6.0 == xs[0]
	expected2 := -4.0 == xs[1]

	if expected != true || expected1 != true || expected2 != true {
		t.Errorf("expected %v, %v, %v got %v, %v, %v", expected, expected1, expected2, len(xs), xs[0], xs[1])
	}
}
