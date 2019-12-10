package main

import (
	"math"
	"os"

	"github.com/kingsleyliao/ray-tracer/src/rendering/light"
	"github.com/kingsleyliao/ray-tracer/src/rendering/material"

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
	drawClockface()
	drawParabola()
	drawRedSphere()
	drawPurpleGiant()
}

func drawPurpleGiant() {
	rayOrigin := point.NewPoint(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 500

	pixelSize := wallSize / float64(canvasPixels)
	wallExtent := wallSize / 2

	c := canvas.NewCanvas(canvasPixels, canvasPixels)

	s := shape.NewSphere()
	s.Material = material.NewMaterial()
	s.Material.Color = color.NewColor(1, 0.2, 1)

	lightPosition := point.NewPoint(-10, 10, -10)
	lightColor := color.NewColor(1, 1, 1)
	l := light.NewPointLight(lightPosition, lightColor)

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
			if hit, ok := intersection.Hit(xs); ok {
				point := r.PositionAt(hit.T)
				normal := hit.Object.NormalAt(point)
				eye := r.Direction.Invert()
				pixelColor := light.Lighting(hit.Object.Material, l, point, eye, normal)
				canvas.WritePixel(c, x, y, pixelColor)
			}
		}
	}

	ppm := c.ToPPM()
	outputPPM(ppm, "purple-giant.ppm")
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
