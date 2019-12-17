package renders

import (
	"math"

	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/src/rendering/canvas"
	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
)

// DrawClockFace draws a clock
func DrawClockFace() {
	width := 500
	height := 500
	c := canvas.NewCanvas(width, height)

	white := color.NewColor(1, 1, 1)

	radius := (3 / 8.0) * float64(width)
	center := width / 2

	position := point.NewPoint(0, 0, 1)
	rotation := matrix.RotationMatrix(vector.Up(), math.Pi/6)
	for i := 0; i < 13; i++ {
		position = rotation.MultiplyTuple(position)
		canvas.WritePixel(c, int(position.X*radius)+center, int(position.Z*radius)+center, white)
	}

	ppm := c.ToPPM()
	outputPPM(ppm, "clockface.ppm")
}
