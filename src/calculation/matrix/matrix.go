package matrix

import (
	"math"

	"github.com/kingsleyliao/ray-tracer/src/calculation/tuple"
)

// Matrix represents the rotation, scalar, and positional data for a node
type Matrix [][]float64

// ZeroMatrix creates a new 4x4 Matrix with zero values
func ZeroMatrix() Matrix {
	return Matrix{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
}

// IdentityMatrix returns a new instance of an identity Matrix
func IdentityMatrix() Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// TranslationMatrix returns a Matrix to be used for translating points. Vectors cannot be translated
func TranslationMatrix(x, y, z float64) Matrix {
	return Matrix{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}

// ScalingMatrix returns a Matrix to be used for scaling points or vectors
func ScalingMatrix(x, y, z float64) Matrix {
	return Matrix{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}

// RotationMatrix returns a Matrix to rotate a point or vector along the specified axis
func RotationMatrix(axis tuple.Tuple, r float64) Matrix {
	if axis.X != 0 {
		return Matrix{
			{1, 0, 0, 0},
			{0, math.Cos(r), -math.Sin(r), 0},
			{0, math.Sin(r), math.Cos(r), 0},
			{0, 0, 0, 1},
		}
	} else if axis.Y != 0 {
		return Matrix{
			{math.Cos(r), 0, math.Sin(r), 0},
			{0, 1, 0, 0},
			{-math.Sin(r), 0, math.Cos(r), 0},
			{0, 0, 0, 1},
		}
	}
	return Matrix{
		{math.Cos(r), -math.Sin(r), 0, 0},
		{math.Sin(r), math.Cos(r), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// ShearMatrix returns a Matrix to be used to shear points
func ShearMatrix(xy, xz, yx, yz, zx, zy float64) Matrix {
	return Matrix{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	}
}

// Multiply returns results the result of Matrix being multiplied by another Matrix
func (m Matrix) Multiply(a Matrix) Matrix {
	result := ZeroMatrix()

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

// Transpose returns the transposed form of a Matrix
func (m Matrix) Transpose() Matrix {
	transposed := ZeroMatrix()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			transposed[i][j] = m[j][i]
		}
	}

	return transposed
}

// Determinant returns the determinant of a 2x2 Matrix
func (m Matrix) Determinant() float64 {
	var d float64
	if len(m) == 2 {
		d = m[0][0]*m[1][1] - m[0][1]*m[1][0]
	} else {
		for i := range m {
			d += m[0][i] * m.Cofactor(0, i)
		}
	}
	return d
}

// SubMatrix returns the subMatrix of a Matrix by ignoring row r and column c
func (m Matrix) SubMatrix(r, c int) Matrix {
	result := Matrix{}
	rowOffset := 0

	for i := 0; i < len(m); i++ {
		if i == r {
			rowOffset = 1
			continue
		}
		row := make([]float64, 0)
		if c == 0 {
			row = append(row, m[i][c+1:]...)
		} else {
			row = append(row, m[i][0:c]...)
			row = append(row, m[i][c+1:]...)
		}

		result = append(result, row)
		for j := range row {
			result[i-rowOffset][j] = row[j]
		}

	}

	return result
}

// Minor returns the minor of a Matrix (the determinant of the subMatrix of a Matrix ignoring row r and column c)
func (m Matrix) Minor(r, c int) float64 {
	sub := m.SubMatrix(r, c)
	return sub.Determinant()
}

// Cofactor returns the cofactor of a Matrix
func (m Matrix) Cofactor(i, j int) float64 {
	minor := m.Minor(i, j)
	if (i+j)%2 == 0 {
		return minor
	}
	return -minor
}

// Equals returns the value equality of two matrices
func (m Matrix) Equals(a Matrix) bool {
	EPSILON := 0.00001
	if len(m) != len(a) {
		return false
	}
	for i := range m {
		for j := range a[i] {
			if math.Abs(m[i][j]-a[i][j]) > EPSILON {
				return false
			}
		}
	}
	return true
}

// IsInvertible returns whether a Matrix can be inverted
func (m Matrix) IsInvertible() bool {
	d := m.Determinant()
	if d == float64(0) {
		return false
	}
	return true
}

// Clone returns a deep copy of a Matrix
func (m Matrix) Clone() Matrix {
	m2 := Matrix{}

	for i := range m {
		m2 = append(m2, make([]float64, len(m[i])))
		for j := range m[i] {
			m2[i][j] = m[i][j]
		}
	}
	return m2
}

// Invert calculates the inverted form of a Matrix
func (m Matrix) Invert() Matrix {
	if !m.IsInvertible() {
		panic("Matrix is not invertible")
	}

	m2 := m.Clone()

	for row := range m {
		for col := range m {
			c := m.Cofactor(row, col)
			// [col, row] here accomplishes the transpose operation
			m2[col][row] = c / m.Determinant()
		}
	}
	return m2
}
