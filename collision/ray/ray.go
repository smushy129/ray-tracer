package ray

import (
	"github.com/kingsleyliao/ray-tracer/calculation/point"
	"github.com/kingsleyliao/ray-tracer/calculation/vector"
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
