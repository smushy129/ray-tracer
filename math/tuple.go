package math

type Tuple struct {
	x float32
	y float32
	z float32
	w int
}

// Point returns a new point in 3D Euclidean space
func Point(x, y, z float32) Tuple {
	return Tuple{x, y, z, 1}
}

// Vector returns a new vector in 3D Euclidean space
func Vector(x, y, z float32) Tuple {
	return Tuple{x, y, z, 0}
}
