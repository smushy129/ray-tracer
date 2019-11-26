package ray

import (
	"github.com/kingsleyliao/ray-tracer/calculation/point"
	"github.com/kingsleyliao/ray-tracer/calculation/vector"
)

// Ray represents a vector that travels through space and hits objects
type Ray struct {
	origin    point.Point
	direction vector.Vector
}

// NewRay creates a new ray
func NewRay(origin point.Point, direction vector.Vector) Ray {
	return Ray{
		origin,
		direction,
	}
}

// PositionAt calculates the position of a ray at time T
func (r Ray) PositionAt(time float64) point.Point {
	return r.origin.Add(r.direction.Scale(time))
}
