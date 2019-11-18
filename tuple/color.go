package tuple

// Color is a tuple RGB values
type Color = Tuple

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