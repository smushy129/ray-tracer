package matrix

import "github.com/kingsleyliao/ray-tracer/tuple"

// Matrix represents the rotation, scalar, and positional data for a node
type Matrix [][]float64

// ZeroMatrix creates a new 4x4 matrix with zero values
func ZeroMatrix() Matrix {
	return Matrix{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
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

// Transpose returns the transposed form of a matrix
func (m Matrix) Transpose() Matrix {
	transposed := ZeroMatrix()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			transposed[i][j] = m[j][i]
		}
	}

	return transposed
}

// Determinant returns the determinant of a 2x2 matrix
func (m Matrix) Determinant() float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

// SubMatrix returns the submatrix of a matrix by ignoring row r and column c
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
			row = append(m[i][0:c], m[i][c+1:]...)
		}

		result = append(result, row)
		for j := range row {
			result[i-rowOffset][j] = row[j]
		}

	}

	return result
}

// Minor returns the minor of a matrix (the determinant of the submatrix of a matrix ignoring row r and column c)
func (m Matrix) Minor(r, c int) float64 {
	sub := m.SubMatrix(r, c)
	return sub.Determinant()
}

// Cofactor returns the cofactor of a matrix
func (m Matrix) Cofactor(i, j int) float64 {
	minor := m.Minor(i, j)
	if (i+j)%2 == 0 {
		return minor
	}
	return -minor
}

// IsEqual returns the value equality of two matrices
func (m Matrix) IsEqual(a Matrix) bool {
	if len(m) != len(a) {
		return false
	}
	for i := range m {
		for j := range a[i] {
			if m[i][j] != a[i][j] {
				return false
			}
		}
	}
	return true
}
