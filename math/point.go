package math

// Point returns a new point in 3D Euclidean space
func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}
