package math

// Vector returns a new vector in 3D Euclidean space
func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}
