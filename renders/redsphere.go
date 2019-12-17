package renders

import (
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/collision/intersection"
	"github.com/kingsleyliao/ray-tracer/src/collision/ray"
	"github.com/kingsleyliao/ray-tracer/src/rendering/canvas"
	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
	"github.com/kingsleyliao/ray-tracer/src/rendering/shape"
)

// DrawRedGiant draws a red giant
func DrawRedGiant() {
	rayOrigin := point.NewPoint(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 500

	pixelSize := wallSize / float64(canvasPixels)
	wallExtent := wallSize / 2

	c := canvas.NewCanvas(canvasPixels, canvasPixels)
	color := color.NewColor(1, 0, 0)
	s := shape.NewSphere()

	// For each row of pixels on the canvas
	for y := 0; y < canvasPixels; y++ {

		// Compute the world Y coordinate (top = +wallExtent, bottom = -wallExtent)
		worldY := wallExtent - pixelSize*float64(y)

		// For each column of pixels
		for x := 0; x < canvasPixels; x++ {

			// Compute the world Y coordinate (left = -wallExtent, right = +wallExtent)
			worldX := -wallExtent + pixelSize*float64(x)

			// Describes the point on the wall that the ray will target
			position := point.NewPoint(worldX, worldY, wallZ)

			r := ray.NewRay(rayOrigin, position.Subtract(rayOrigin).Normalize())
			xs := intersection.Intersect(s, r)

			// _, ok := intersection.Hit(xs)
			// fmt.Println(ok)
			if _, ok := intersection.Hit(xs); ok {
				canvas.WritePixel(c, x, y, color)
			}
		}
	}

	ppm := c.ToPPM()
	outputPPM(ppm, "redgiant.ppm")
}
