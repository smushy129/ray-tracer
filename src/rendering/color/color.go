package color

import "math"

// Color is a Color RGB values
type Color struct {
	R float64
	G float64
	B float64
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
		c.R*c2.R,
		c.G*c2.G,
		c.B*c2.B,
	)
}

// Add adds two Colors
func (c Color) Add(c1 Color) Color {
	return Color{
		R: c.R + c1.R,
		G: c.G + c1.G,
		B: c.B + c1.B,
		W: c.W + c1.W,
	}
}

// Subtract subtracts two Colors
func (c Color) Subtract(c1 Color) Color {
	return Color{
		R: c.R - c1.R,
		G: c.G - c1.G,
		B: c.B - c1.B,
		W: c.W - c1.W,
	}
}

// Scale multiplies a scalar value by each value of a Color
func (c Color) Scale(s float64) Color {
	return Color{
		R: c.R * s,
		G: c.G * s,
		B: c.B * s,
		W: c.W * s,
	}
}

// Equals compares Colors of equality
func (c Color) Equals(c1 Color) bool {
	EPSILON := 0.00001
	if math.Abs(c.R-c1.R) > EPSILON {
		return false
	}
	if math.Abs(c.G-c1.G) > EPSILON {
		return false
	}
	if math.Abs(c.B-c1.B) > EPSILON {
		return false
	}
	if c.W != c1.W {
		return false
	}
	return true
}
