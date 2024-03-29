package shape

import (
	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/src/rendering/material"
)

// Sphere represents a unit sphere
type Sphere struct {
	Center    point.Point
	Radius    vector.Vector
	Transform matrix.Matrix
	Material  material.Material
}

// NewSphere returns a Sphere
func NewSphere() Sphere {
	return Sphere{
		Center:    point.NewPoint(0, 0, 0),
		Radius:    vector.NewVector(1, 1, 1),
		Transform: matrix.IdentityMatrix(),
		Material:  material.DefaultMaterial(),
	}
}

// SetTransform mutates the Transform Matrix of a Sphere
func (s *Sphere) SetTransform(m matrix.Matrix) {
	s.Transform = m
}

// SetMaterial mutates the Transform Matrix of a Sphere
func (s *Sphere) SetMaterial(m material.Material) {
	s.Material = m
}

// NormalAt finds the normal vector of a point on the surface of a Sphere
func (s Sphere) NormalAt(p point.Point) vector.Vector {
	objectPoint := s.Transform.Invert().MultiplyTuple(p)
	objectNormal := objectPoint.Subtract(point.Zero())
	worldNormal := s.Transform.Invert().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}

// Equals compares Spheres for equality
func (s Sphere) Equals(s2 Sphere) bool {
	return s.Center == s2.Center &&
		s.Radius == s2.Radius &&
		s.Transform.Equals(s2.Transform)
}
