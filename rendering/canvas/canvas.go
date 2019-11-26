package canvas

import (
	"fmt"
	"math"
	"strings"

	"github.com/kingsleyliao/ray-tracer/rendering/color"
)

// Canvas is for drawing pixels on
type Canvas struct {
	Width  int
	Height int
	Matrix [][]color.Color
}

// NewCanvas is a constructor for a Canvas
func NewCanvas(w, h int) Canvas {
	m := make([][]color.Color, h)
	for i := range m {
		m[i] = make([]color.Color, w)
	}
	return Canvas{
		w,
		h,
		m,
	}
}

// ToPPM generates the string necessary to save the cavas as a PPM file
func (c Canvas) ToPPM() string {
	var rows string

	for i := range c.Matrix {
		row := c.Matrix[i]
		line := ""
		for _, color := range row {
			line += fmt.Sprintf("%d ", clamp(0, 255, color.X*255))
			line += fmt.Sprintf("%d ", clamp(0, 255, color.Y*255))
			line += fmt.Sprintf("%d ", clamp(0, 255, color.Z*255))
		}
		line = strings.Trim(line, " ") + "\n"
		rows += line
	}
	header := createPpmHeader(c)
	ppm := header + rows
	return ppm
}

func createPpmHeader(c Canvas) string {
	return fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height)
}

// WritePixel writes a color to a canvas
func WritePixel(c Canvas, width, height int, color color.Color) error {
	if isOutOfBounds(c, width, height) {
		return fmt.Errorf("%s", "Dimensions out of bounds")
	}
	c.Matrix[height][width] = color
	return nil
}

// PixelAt returns the color of a pixel for given xy coordinates
func PixelAt(c Canvas, width, height int) (color.Color, error) {
	if isOutOfBounds(c, width, height) {
		return color.NewColor(0, 0, 0), fmt.Errorf("%s", "Dimensions out of bounds")
	}
	return c.Matrix[height][width], nil
}

func isOutOfBounds(c Canvas, x, y int) bool {
	if x >= c.Width || y >= c.Height || x < 0 || y < 0 {
		return true
	}
	return false
}

func clamp(min, max, x float64) int {
	return int(math.Round(math.Min(math.Max(x, min), max)))
}
