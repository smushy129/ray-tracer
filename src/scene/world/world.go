package world

import (
	"sort"

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

// World represents the entire 3D scene
type World struct {
	Objects []shape.Sphere
	Light   light.Light
}

// DefaultWorld creates a World with 1 point light and 2 Spheres
func DefaultWorld() World {
	s1 := shape.NewSphere()
	s1.Material = material.Material{
		Color:     color.NewColor(0.8, 1.0, 0.6),
		Ambient:   0.1,
		Diffuse:   0.7,
		Specular:  0.2,
		Shininess: 200.0,
	}
	s2 := shape.NewSphere()
	s2.Transform = matrix.ScalingMatrix(0.5, 0.5, 0.5)

	l := light.NewPointLight(point.NewPoint(-10, 10, -10), color.NewColor(1, 1, 1))

	return World{
		Objects: []shape.Sphere{s1, s2},
		Light:   l,
	}
}

// IntersectWorld returns a all the objects in a world intersected by a Ray
func IntersectWorld(w World, r ray.Ray) intersection.ByT {
	intersections := intersection.ByT{}
	for _, s := range w.Objects {
		intersections = append(intersections, intersection.Intersect(s, r)...)
	}
	sort.Sort(intersections)
	return intersections
}

// ShadingComputations are computations used for shading effectively
type ShadingComputations struct {
	T         float64
	Object    shape.Sphere
	Point     point.Point
	EyeV      vector.Vector
	NormalV   vector.Vector
	Inside    bool
	OverPoint point.Point
}

// PrepareShadeHit the state of an Intersection
func PrepareShadeHit(i intersection.Intersection, r ray.Ray) ShadingComputations {
	EPSILON := 0.00001
	eyeV := r.Direction.Invert()
	normalV := i.Object.NormalAt(r.PositionAt(i.T))
	inside := false
	p := r.PositionAt(i.T)
	overPoint := p.Add(normalV.Scale(EPSILON))
	// Negative dot product of two vectors means the vectors are pointing in opposite directions
	// and that the ray originates from inside the object, behind the intersection
	if normalV.Dot(eyeV) < 0 {
		inside = true
		normalV = normalV.Invert()
	}
	return ShadingComputations{
		T:         i.T,
		Object:    i.Object,
		Point:     p,
		EyeV:      eyeV,
		NormalV:   normalV,
		Inside:    inside,
		OverPoint: overPoint,
	}
}

// ShadeHit shades the color of a pixel based no the properties of the Ray and Object
func (w World) ShadeHit(comps ShadingComputations) color.Color {
	shadowed := w.IsShadowed(comps.OverPoint)
	return light.Lighting(
		comps.Object.Material,
		w.Light,
		comps.Point,
		comps.EyeV,
		comps.NormalV,
		shadowed,
	)
}

// ColorAt returns the color at a position in a world
func (w World) ColorAt(r ray.Ray) color.Color {
	xs := IntersectWorld(w, r)
	hit, ok := intersection.Hit(xs)
	if !ok {
		return color.Black()
	}
	comps := PrepareShadeHit(hit, r)
	return w.ShadeHit(comps)
}

// ViewTransform describes the World's default orientation
func ViewTransform(from, to point.Point, up vector.Vector) matrix.Matrix {
	forward := to.Subtract(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)

	orientation := matrix.Matrix{
		{left.X, left.Y, left.Z, 0},
		{trueUp.X, trueUp.Y, trueUp.Z, 0},
		{-forward.X, -forward.Y, -forward.Z, 0},
		{0, 0, 0, 1},
	}
	return orientation.Multiply(matrix.TranslationMatrix(-from.X, -from.Y, -from.Z))
}

// IsShadowed returns if a pixel is in a shadow
func (w World) IsShadowed(p point.Point) bool {
	v := w.Light.Position.Subtract(p)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := ray.NewRay(p, direction)
	xs := IntersectWorld(w, r)

	h, ok := intersection.Hit(xs)

	if !ok {
		return false
	} else if h.T < distance {
		return true
	}
	return false
}
