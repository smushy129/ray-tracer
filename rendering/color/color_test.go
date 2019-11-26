package color

import "testing"

func TestNewColor(t *testing.T) {
	c1 := NewColor(-0.5, 0.4, 1.7)
	c2 := NewColor(-0.5, 0.4, 1.7)

	if c2.X != c1.X ||
		c2.Y != c1.Y ||
		c2.Z != c1.Z ||
		c2.W != c1.W {
		t.Errorf("got %v, expected %v", c2, c1)
	}
}

func TestAddColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	result := c1.Add(c2)
	expected := NewColor(1.6, 0.7, 1)

	if !result.Equals(expected) {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

func TestSubtractColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	result := c1.Subtract(c2)
	expected := NewColor(0.2, 0.5, 0.5)

	if !result.Equals(expected) {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

func TestScaleColors(t *testing.T) {
	c1 := NewColor(0.2, 0.3, 0.4)

	result := c1.Scale(2)
	expected := NewColor(0.4, 0.6, 0.8)

	if !result.Equals(expected) {
		t.Errorf("got %v, expected %v", result, expected)
	}
}

func TestMultiplyColors(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)

	result := c1.Multiply(c2)
	expected := NewColor(0.9, 0.2, 0.04)

	if !result.Equals(expected) {
		t.Errorf("got %v, expected %v", result, expected)
	}
}
