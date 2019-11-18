package main

import (
	"os"

	"github.com/kingsleyliao/ray-tracer/canvas"
	"github.com/kingsleyliao/ray-tracer/projectile"
	"github.com/kingsleyliao/ray-tracer/tuple"
)

func main() {
	c := canvas.NewCanvas(900, 550)
	red := tuple.NewColor(1, 0, 0)

	gravity := tuple.NewVector(0, -0.1, 0)
	drag := tuple.NewVector(-0.01, 0, 0)
	env := projectile.NewEarth(gravity, drag)

	velocity := tuple.NewVector(1, 1.8, 0).Normalize().Scale(11.25)
	start := tuple.NewPoint(0, 1, 0)
	proj := projectile.NewRay(start, velocity)

	for proj.Point.Y > 0 {
		canvas.WritePixel(c, int(proj.Point.X), 550-int(proj.Point.Y), red)
		proj = proj.Tick(env)
	}

	ppm := c.ToPPM()

	f, err := os.Create("canvas.ppm")
	if err != nil {
		panic(err)
	}

	f.WriteString(ppm)

}
