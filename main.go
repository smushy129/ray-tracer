package main

import (
	"fmt"

	"github.com/kingsleyliao/ray-tracer/matrix"
)

// func main() {
// 	c := canvas.NewCanvas(900, 550)
// 	red := tuple.NewColor(1, 0, 0)

// 	gravity := tuple.NewVector(0, -0.1, 0)
// 	drag := tuple.NewVector(-0.01, 0, 0)
// 	env := projectile.NewEarth(gravity, drag)

// 	velocity := tuple.NewVector(1, 1.8, 0).Normalize().Scale(11.25)
// 	start := tuple.NewPoint(0, 1, 0)
// 	proj := projectile.NewRay(start, velocity)

// 	for proj.Point.Y > 0 {
// 		canvas.WritePixel(c, int(proj.Point.X), 550-int(proj.Point.Y), red)
// 		proj = proj.Tick(env)
// 	}

// 	ppm := c.ToPPM()

// 	f, err := os.Create("canvas.ppm")
// 	if err != nil {
// 		panic(err)
// 	}

// 	f.WriteString(ppm)

// }

func main() {
	m := matrix.Matrix{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	}

	m2 := m.Multiply(m.Invert())
	printS(m2)
}

func printS(s [][]float64) {
	for i := range s {
		fmt.Println(s[i])
	}
}
