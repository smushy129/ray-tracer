package tuple

import "math"

// Tuple represents Points and Vectors
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Equals compares float for equality within a range
func (t Tuple) Equals(t1 Tuple) bool {
	EPSILON := 0.00001
	if math.Abs(t.X-t1.X) > EPSILON {
		return false
	}
	if math.Abs(t.Y-t1.Y) > EPSILON {
		return false
	}
	if math.Abs(t.Z-t1.Z) > EPSILON {
		return false
	}
	if t.W != t1.W {
		return false
	}
	return true
}

// Add adds two tuples
func (t Tuple) Add(t1 Tuple) Tuple {
	return Tuple{
		X: t.X + t1.X,
		Y: t.Y + t1.Y,
		Z: t.Z + t1.Z,
		W: t.W + t1.W,
	}
}

// Subtract subtracts two tuples
func (t Tuple) Subtract(t1 Tuple) Tuple {
	return Tuple{
		X: t.X - t1.X,
		Y: t.Y - t1.Y,
		Z: t.Z - t1.Z,
		W: t.W - t1.W,
	}
}

// Scale multiplies a scalar value by each value of a tuple
func (t Tuple) Scale(s float64) Tuple {
	return Tuple{
		X: t.X * s,
		Y: t.Y * s,
		Z: t.Z * s,
		W: t.W * s,
	}
}

// Invert negates each value of a tuple
func (t Tuple) Invert() Tuple {
	return Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}
}
