package matrix

import "github.com/kingsleyliao/ray-tracer/tuple"

// Matrix represents the rotation, scalar, and positional data for a node
type Matrix [4][4]float64

// ZeroMatrix creates a new 4x4 matrix with zero values
func ZeroMatrix() Matrix {
	return Matrix{}
}

// IdentityMatrix returns a new instance of an identity matrix
func IdentityMatrix() Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
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

// Transpose returns the transposed form of a matrix
func (m Matrix) Transpose() Matrix {
	transposed := Matrix{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			transposed[i][j] = m[j][i]
		}
	}

	return transposed
}

// Determinant returns the determinant of a 2x2 matrix
func Determinant(m [2][2]float64) float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

// SubMatrix returns the submatrix of a matrix by ignoring row r and column c
func SubMatrix(m Matrix, r, c int) [][]float64 {
	result := [][]float64{}
	rowOffset := 0

	for i := 0; i < len(m); i++ {
		if i == r {
			rowOffset = 1
			continue
		}
		row := append(m[i][0:c], m[i][c+1:]...)
		result = append(result, row)
		for j := range row {
			result[i-rowOffset][j] = row[j]
		}

	}

	return result
}
