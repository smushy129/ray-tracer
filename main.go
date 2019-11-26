package main

import (
	"fmt"
	"math"
	"os"

	"github.com/kingsleyliao/ray-tracer/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/calculation/point"
	"github.com/kingsleyliao/ray-tracer/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/collision/projectile"
	"github.com/kingsleyliao/ray-tracer/rendering/canvas"
	"github.com/kingsleyliao/ray-tracer/rendering/color"
)

const rendersDir = "./renders/"

func main() {
	drawClockface()
	drawParabola()
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

	f, err := os.Create(rendersDir + "clock-face.ppm")
	if err != nil {
		panic(err)
	}

	f.WriteString(ppm)

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

	f, err := os.Create(rendersDir + "inverted-parabola.ppm")
	if err != nil {
		panic(err)
	}

	f.WriteString(ppm)
}

func printMatrix(s [][]float64) {
	for i := range s {
		fmt.Println(s[i])
	}
}
