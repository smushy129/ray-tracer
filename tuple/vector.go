package tuple

import "math"

// Vector is a type alias for Tuple with w = 0
type Vector = Tuple

// NewVector returns a new vector in 3D Euclidean space
func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z, 0}
}

// Normalize normalizes each value of a Vector
func (v Vector) Normalize() Vector {
	m := v.Magnitude()
	return Vector{
		v.X / m,
		v.Y / m,
		v.Z / m,
		v.W / m,
	}
}

// Magnitude finds the distance a vector represents with respect to the origin
func (v Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2) + math.Pow(v.W, 2))
}

// Dot computes the dot product between two vectors
func (v Vector) Dot(b Vector) float64 {
	return v.X*b.X +
		v.Y*b.Y +
		v.Z*b.Z +
		v.W*b.W
}

// Cross computes the cross product between two vectors
func (v Vector) Cross(b Vector) Vector {
	return NewVector(
		v.Y*b.Z-v.Z*b.Y,
		v.Z*b.X-v.X*b.Z,
		v.X*b.Y-v.Y*b.X,
	)
}
