package math

import "math"

// Vector is a type alias for Tuple with w = 0
type Vector = Tuple

// NewVector returns a new vector in 3D Euclidean space
func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z, 0}
}

func (v Vector) normalize() Vector {
	m := v.magnitude()
	return Vector{
		v.x / m,
		v.y / m,
		v.z / m,
		v.w / m,
	}
}

func (v Vector) magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2) + math.Pow(v.w, 2))
}

func (v Vector) dot(b Vector) float64 {
	return v.x*b.x +
		v.y*b.y +
		v.z*b.z +
		v.w*b.w
}

func (v Vector) cross(b Vector) Vector {
	return NewVector(
		v.y*b.z-v.z*b.y,
		v.z*b.x-v.x*b.z,
		v.x*b.y-v.y*b.x,
	)
}
