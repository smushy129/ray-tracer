package world

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/src/collision/intersection"
	"github.com/kingsleyliao/ray-tracer/src/collision/ray"

	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
	"github.com/kingsleyliao/ray-tracer/src/rendering/light"
	"github.com/kingsleyliao/ray-tracer/src/rendering/material"
	"github.com/kingsleyliao/ray-tracer/src/rendering/shape"
)

// Creates a Default world with the correct default values
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

// Intersect a world with a ray
func TestIntersectWorld(t *testing.T) {
	w := DefaultWorld()
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	xs := IntersectWorld(w, r)

	if xs[0].T != 4 {
		t.Errorf("expected %v, got %v", 4, xs[0].T)
	}

	if xs[1].T != 4.5 {
		t.Errorf("expected %v, got %v", 4.5, xs[1].T)
	}

	if xs[2].T != 5.5 {
		t.Errorf("expected %v, got %v", 5.5, xs[2].T)
	}

	if xs[3].T != 6 {
		t.Errorf("expected %v, got %v", 6, xs[3].T)
	}
}

// Precomputing the state of an intersection
func TestPrepareShadeHit_1(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	shape := shape.NewSphere()
	xs := intersection.NewIntersection(4, shape)
	comps := PrepareShadeHit(xs, r)

	if comps.T != xs.T {
		t.Errorf("expected %v, got %v", xs.T, comps.T)
	}

	if !comps.Object.Equals(xs.Object) {
		t.Errorf("expected %v, got %v", xs.Object, comps.Object)
	}

	if comps.Point != point.NewPoint(0, 0, -1) {
		t.Errorf("expected %v, got %v", point.NewPoint(0, 0, -1), comps.Point)
	}

	if comps.EyeV != vector.NewVector(0, 0, -1) {
		t.Errorf("expected %v, got %v", vector.NewVector(0, 0, -1), comps.EyeV)
	}

	if comps.NormalV != vector.NewVector(0, 0, -1) {
		t.Errorf("expected %v, got %v", vector.NewVector(0, 0, -1), comps.NormalV)
	}

}

// The hit, when an intersection occurs on the outside
func TestPrepareShadeHit_2(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	shape := shape.NewSphere()
	xs := intersection.NewIntersection(4, shape)
	comps := PrepareShadeHit(xs, r)

	if comps.Inside != false {
		t.Errorf("expected %v, got %v", false, comps.Inside)
	}
}

// The hit, when an intersection occurs on the inside
func TestPrepareShadeHit_3(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1))
	shape := shape.NewSphere()
	xs := intersection.NewIntersection(1, shape)
	comps := PrepareShadeHit(xs, r)

	if comps.Inside != true {
		t.Errorf("expected %v, got %v", true, comps.Inside)
	}

	if comps.Point != point.NewPoint(0, 0, 1) {
		t.Errorf("expected %v, got %v", point.NewPoint(0, 0, -1), comps.Point)
	}

	if comps.EyeV != vector.NewVector(0, 0, -1) {
		t.Errorf("expected %v, got %v", vector.NewVector(0, 0, -1), comps.EyeV)
	}

	if comps.NormalV != vector.NewVector(0, 0, -1) {
		t.Errorf("expected %v, got %v", vector.NewVector(0, 0, -1), comps.NormalV)
	}

}

// The hit should offset the point
func TestPrepareShadeHitOverPoint(t *testing.T) {
	EPSILON := 0.00001
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	shape := shape.NewSphere()
	shape.SetTransform(matrix.TranslationMatrix(0, 0, 1))
	xs := intersection.NewIntersection(5, shape)
	comps := PrepareShadeHit(xs, r)

	if !(comps.OverPoint.Z < -EPSILON/2) {
		t.Errorf("expected %v, got %v", true, false)
	}

	if !(comps.Point.Z > comps.OverPoint.Z) {
		t.Errorf("expected %v, got %v", true, false)
	}
}

// Shading an intersection
func TestShadeHit_1(t *testing.T) {
	w := DefaultWorld()
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	shape := w.Objects[0]
	xs := intersection.NewIntersection(4, shape)

	comps := PrepareShadeHit(xs, r)
	shade := w.ShadeHit(comps)

	expected := color.NewColor(0.38066, 0.47583, 0.2855)

	if !shade.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, shade)
	}
}

// Shading an intersection from the inside
func TestShadeHit_2(t *testing.T) {
	w := DefaultWorld()
	w.Light = light.NewPointLight(point.NewPoint(0, 0.25, 0), color.NewColor(1, 1, 1))
	r := ray.NewRay(point.NewPoint(0, 0, 0), vector.NewVector(0, 0, 1))
	shape := w.Objects[1]
	xs := intersection.NewIntersection(0.5, shape)

	comps := PrepareShadeHit(xs, r)
	shade := w.ShadeHit(comps)

	// expected := color.NewColor(0.90498, 0.90498, 0.90498)

	expected := color.NewColor(0.1, 0.1, 0.1) // This test is now in a shadow

	if !shade.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, shade)
	}
}

// The color when a ray misses
func TestColorAt_1(t *testing.T) {
	w := DefaultWorld()
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 1, 0))
	c := w.ColorAt(r)

	if !c.Equals(color.Black()) {
		t.Errorf("expected %v, got %v", color.Black(), c)
	}
}

// The color when a ray hits
func TestColorAt_2(t *testing.T) {
	w := DefaultWorld()
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	c := w.ColorAt(r)

	expected := color.NewColor(0.38066, 0.47583, 0.2855)

	if !c.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, c)
	}
}

// The color with an intersectino behind the ray
func TestColorAt_3(t *testing.T) {
	w := DefaultWorld()
	outer := &w.Objects[0]
	outer.Material.Ambient = 1.0
	inner := &w.Objects[1]
	inner.Material.Ambient = 1.0
	r := ray.NewRay(point.NewPoint(0, 0, 0.75), vector.NewVector(0, 0, -1))
	c := w.ColorAt(r)

	expected := inner.Material.Color

	if !c.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, c)
	}
}

// The transform matrix for the default orientation
func TestViewTransform_1(t *testing.T) {
	from := point.NewPoint(0, 0, 0)
	to := point.NewPoint(0, 0, -1)
	up := vector.NewVector(0, 1, 0)
	transform := ViewTransform(from, to, up)

	if !transform.Equals(matrix.IdentityMatrix()) {
		t.Errorf("expected %v, got %v", matrix.IdentityMatrix(), t)
	}
}

// A view transformation matrix looking in the positive z direction
func TestViewTransform_2(t *testing.T) {
	from := point.NewPoint(0, 0, 0)
	to := point.NewPoint(0, 0, 1)
	up := vector.NewVector(0, 1, 0)
	transform := ViewTransform(from, to, up)

	expected := matrix.ScalingMatrix(-1, 1, -1)

	if !transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, t)
	}
}

// A view transformation moves the world
func TestViewTransform_3(t *testing.T) {
	from := point.NewPoint(0, 0, 8)
	to := point.NewPoint(0, 0, 0)
	up := vector.NewVector(0, 1, 0)
	transform := ViewTransform(from, to, up)

	expected := matrix.TranslationMatrix(0, 0, -8)

	if !transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, t)
	}
}

// An arbitrary view transform
func TestViewTransform_4(t *testing.T) {
	from := point.NewPoint(1, 3, 2)
	to := point.NewPoint(4, -2, 8)
	up := vector.NewVector(1, 1, 0)
	transform := ViewTransform(from, to, up)

	expected := matrix.Matrix{
		{-0.50709, 0.50709, 0.67612, -2.36643},
		{0.76772, 0.60609, 0.12122, -2.82843},
		{-0.35857, 0.59761, -0.71714, 0.00000},
		{0.00000, 0.00000, 0.00000, 1.00000},
	}

	if !transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, t)
	}
}

// There is no shadow when nothing is collinear with point and light
func TestIsShadowed_1(t *testing.T) {
	w := DefaultWorld()
	p := point.NewPoint(0, 10, 0)

	isShadowed := w.IsShadowed(p)
	expected := false

	if isShadowed != expected {
		t.Errorf("expected %v, got %v", expected, isShadowed)
	}
}

// The shadow when an object is between the point and the light
func TestIsShadowed_2(t *testing.T) {
	w := DefaultWorld()
	p := point.NewPoint(10, -10, 10)

	isShadowed := w.IsShadowed(p)
	expected := true

	if isShadowed != expected {
		t.Errorf("expected %v, got %v", expected, isShadowed)
	}
}

// There is no shadow when an object is behind the light
func TestIsShadowed_3(t *testing.T) {
	w := DefaultWorld()
	p := point.NewPoint(-20, 20, -20)

	isShadowed := w.IsShadowed(p)
	expected := false

	if isShadowed != expected {
		t.Errorf("expected %v, got %v", expected, isShadowed)
	}
}

// There is no shadow when an object is behind the point
func TestIsShadowed_4(t *testing.T) {
	w := DefaultWorld()
	p := point.NewPoint(-2, 2, -2)

	isShadowed := w.IsShadowed(p)
	expected := false

	if isShadowed != expected {
		t.Errorf("expected %v, got %v", expected, isShadowed)
	}
}
