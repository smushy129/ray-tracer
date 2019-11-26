package point

import "github.com/kingsleyliao/ray-tracer/calculation/tuple"

// Point is a Tuple with w = 1
type Point = tuple.Tuple

// NewPoint returns a new point in 3D Euclidean space
func NewPoint(x, y, z float64) Point {
	return Point{X: x, Y: y, Z: z, W: 1}
}
