package tuple

import "math"

// Tuple represents Points and Tuples
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

// Normalize normalizes each value of a Vector
func (t Tuple) Normalize() Tuple {
	m := t.Magnitude()
	return Tuple{
		t.X / m,
		t.Y / m,
		t.Z / m,
		t.W / m,
	}
}

// Magnitude finds the distance a Tuple represents with respect to the origin
func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2))
}

// Dot computes the dot product between two Tuples
func (t Tuple) Dot(b Tuple) float64 {
	return t.X*b.X +
		t.Y*b.Y +
		t.Z*b.Z +
		t.W*b.W
}

// Cross computes the cross product between two Tuples
func (t Tuple) Cross(b Tuple) Tuple {
	return Tuple{
		X: t.Y*b.Z - t.Z*b.Y,
		Y: t.Z*b.X - t.X*b.Z,
		Z: t.X*b.Y - t.Y*b.X,
	}
}

// Reflect returns the reflection vector baesd on the normal of a surface
func (t Tuple) Reflect(normal Tuple) Tuple {
	dot := t.Dot(normal)
	return t.Subtract(normal.Scale(2).Scale(dot))
}
