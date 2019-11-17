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
		v.x / m,
		v.y / m,
		v.z / m,
		v.w / m,
	}
}

// Magnitude finds the distance a vector represents with respect to the origin
func (v Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2) + math.Pow(v.w, 2))
}

// Dot computes the dot product between two vectors
func (v Vector) Dot(b Vector) float64 {
	return v.x*b.x +
		v.y*b.y +
		v.z*b.z +
		v.w*b.w
}

// Cross computes the cross product between two vectors
func (v Vector) Cross(b Vector) Vector {
	return NewVector(
		v.y*b.z-v.z*b.y,
		v.z*b.x-v.x*b.z,
		v.x*b.y-v.y*b.x,
	)
}
