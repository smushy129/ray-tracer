package vector

import (
	"github.com/kingsleyliao/ray-tracer/src/calculation/tuple"
)

// Vector is a type alias for Tuple with w = 0
type Vector = tuple.Tuple

// NewVector returns a new vector in 3D Euclidean space
func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z, W: 0}
}

// Zero returns a Vector of zero values
func Zero() Vector {
	return NewVector(0, 0, 0)
}

// Left returns a Vector pointing in the negative X-axis
func Left() Vector {
	return NewVector(-1, 0, 0)
}

// Right returns a Vector pointing in the positive X-axis
func Right() Vector {
	return NewVector(1, 0, 0)
}

// Back returns a Vector pointing in the positive Z-axis
func Back() Vector {
	return NewVector(0, 0, 1)
}

// Front returns a Vector pointing in the negative Z-axis
func Front() Vector {
	return NewVector(0, 0, -1)
}

// Up returns a Vector pointing in the positve Y-axis
func Up() Vector {
	return NewVector(0, 1, 0)
}

// Down returns a Vector pointing in the negative Y-axis
func Down() Vector {
	return NewVector(0, -1, 0)
}
