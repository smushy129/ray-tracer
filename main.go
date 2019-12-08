package main

import (
	"math"
	"os"

	"github.com/kingsleyliao/ray-tracer/src/collision/intersection"
	"github.com/kingsleyliao/ray-tracer/src/collision/ray"

	"github.com/kingsleyliao/ray-tracer/src/rendering/shape"

	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/src/collision/projectile"
	"github.com/kingsleyliao/ray-tracer/src/rendering/canvas"
	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
)

const rendersDir = "./renders/"

func main() {
	// drawClockface()
	// drawParabola()
	drawRedSphere()
}

func drawRedSphere() {
	rayOrigin := point.NewPoint(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 500

	pixelSize := wallSize / float64(canvasPixels)
	wallExtent := wallSize / 2

	c := canvas.NewCanvas(canvasPixels, canvasPixels)
	color := color.NewColor(1, 0, 0)
	shape := shape.NewSphere()

	for y := 0; y < canvasPixels; y++ {
		worldY := wallExtent - pixelSize*float64(y)

		for x := 0; x < canvasPixels; x++ {
			worldX := -wallExtent + pixelSize*float64(x)

			position := point.NewPoint(worldX, worldY, wallZ)

			r := ray.NewRay(rayOrigin, position.Subtract(rayOrigin).Normalize())
			xs := intersection.Intersect(shape, r)

			// _, ok := intersection.Hit(xs)
			// fmt.Println(ok)
			if _, ok := intersection.Hit(xs); ok {
				canvas.WritePixel(c, x, y, color)
			}
		}
	}

	ppm := c.ToPPM()
	outputPPM(ppm, "red-giant.ppm")
}

func drawClockface() {
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
	outputPPM(ppm, "clock-face.ppm")
}

func drawParabola() {
	c := canvas.NewCanvas(900, 550)
	red := color.NewColor(1, 0, 0)

	gravity := vector.NewVector(0, -0.1, 0)
	drag := vector.NewVector(-0.01, 0, 0)
	env := projectile.NewEarth(gravity, drag)

	velocity := vector.NewVector(1, 1.8, 0).Normalize().Scale(11.25)
	start := point.NewPoint(0, 1, 0)
	proj := projectile.NewRay(start, velocity)

	for proj.Point.Y > 0 {
		canvas.WritePixel(c, int(proj.Point.X), 550-int(proj.Point.Y), red)
		proj = proj.Tick(env)
	}

	ppm := c.ToPPM()
	outputPPM(ppm, "inverted-parabola.ppm")
}

func outputPPM(ppm string, name string) {
	f, err := os.Create(rendersDir + name)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	f.WriteString(ppm)
}
