package tuple

// Point is a Tuple with w = 1
type Point = Tuple

// NewPoint returns a new point in 3D Euclidean space
func NewPoint(x, y, z float64) Point {
	return Point{x, y, z, 1}
}
