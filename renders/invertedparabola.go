package renders

import (
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/src/collision/projectile"
	"github.com/kingsleyliao/ray-tracer/src/rendering/canvas"
	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
)

// DrawParabola draws a parabola
func DrawParabola() {
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
	outputPPM(ppm, "invertedparabola.ppm")
}
