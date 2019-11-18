package matrix

import "github.com/kingsleyliao/ray-tracer/tuple"

// Matrix represents the rotation, scalar, and positional data for a node
type Matrix [4][4]float64

// ZeroMatrix creates a new 4x4 matrix with zero values
func ZeroMatrix() Matrix {
	return Matrix{}
}

// Multiply returns results the result of Matrix being multiplied by another Matrix
func (m Matrix) Multiply(a Matrix) Matrix {
	result := Matrix{}

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			result[row][col] =
				m[row][0]*a[0][col] +
					m[row][1]*a[1][col] +
					m[row][2]*a[2][col] +
					m[row][3]*a[3][col]
		}
	}

	return result
}

// MultiplyTuple multiplies the Matrix by a Tuple
func (m Matrix) MultiplyTuple(t tuple.Tuple) tuple.Tuple {
	result := [4]float64{}
	t1 := [4]float64{t.X, t.Y, t.Z, t.W}
	for row := 0; row < 4; row++ {
		result[row] =
			m[row][0]*t1[0] +
				m[row][1]*t1[1] +
				m[row][2]*t1[2] +
				m[row][3]*t1[3]

	}

	return tuple.Tuple{X: result[0], Y: result[1], Z: result[2], W: result[3]}
}
