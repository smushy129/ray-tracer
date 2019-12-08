package ray

import (
	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
)

// Ray represents a vector that travels through space and hits objects
type Ray struct {
	Origin    point.Point
	Direction vector.Vector
}

// NewRay creates a new ray
func NewRay(origin point.Point, direction vector.Vector) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

// PositionAt calculates the position of a ray at time T
func (r Ray) PositionAt(time float64) point.Point {
	return r.Origin.Add(r.Direction.Scale(time))
}

// Transform returns a new Ray that is translated, rotated, and scaled by a transform matrix
func (r Ray) Transform(m matrix.Matrix) Ray {
	return NewRay(
		m.MultiplyTuple(r.Origin),
		m.MultiplyTuple(r.Direction),
	)
}
