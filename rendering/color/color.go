package color

import "math"

// Color is a Color RGB values
type Color struct {
	X float64
	Y float64
	Z float64
	W float64
}

// NewColor creates a new color
func NewColor(r, g, b float64) Color {
	return Color{
		r,
		g,
		b,
		0,
	}
}

// Multiply multiplies two colors
func (c Color) Multiply(c2 Color) Color {
	return NewColor(
		c.X*c2.X,
		c.Y*c2.Y,
		c.Z*c2.Z,
	)
}

// Add adds two Colors
func (c Color) Add(c1 Color) Color {
	return Color{
		X: c.X + c1.X,
		Y: c.Y + c1.Y,
		Z: c.Z + c1.Z,
		W: c.W + c1.W,
	}
}

// Subtract subtracts two Colors
func (c Color) Subtract(c1 Color) Color {
	return Color{
		X: c.X - c1.X,
		Y: c.Y - c1.Y,
		Z: c.Z - c1.Z,
		W: c.W - c1.W,
	}
}

// Scale multiplies a scalar value by each value of a Color
func (c Color) Scale(s float64) Color {
	return Color{
		X: c.X * s,
		Y: c.Y * s,
		Z: c.Z * s,
		W: c.W * s,
	}
}

// Equals compares Colors of equality
func (c Color) Equals(c1 Color) bool {
	EPSILON := 0.00001
	if math.Abs(c.X-c1.X) > EPSILON {
		return false
	}
	if math.Abs(c.Y-c1.Y) > EPSILON {
		return false
	}
	if math.Abs(c.Z-c1.Z) > EPSILON {
		return false
	}
	if c.W != c1.W {
		return false
	}
	return true
}
