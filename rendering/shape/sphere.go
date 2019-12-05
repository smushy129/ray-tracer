package sphere

import (
	"github.com/kingsleyliao/ray-tracer/calculation/point"
	"github.com/kingsleyliao/ray-tracer/calculation/vector"
)

// Sphere represents a unit sphere
type Sphere struct {
	Center point.Point
	Radius vector.Vector
}

// NewSphere returns a Sphere
func NewSphere() Sphere {
	return Sphere{
		Center: point.NewPoint(0, 0, 0),
		Radius: vector.NewVector(1, 1, 1),
	}
}
