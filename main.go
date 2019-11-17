package main

import (
	"os"

	"github.com/kingsleyliao/ray-tracer/canvas"
	"github.com/kingsleyliao/ray-tracer/tuple"
)

func main() {
	c := canvas.NewCanvas(1280, 720)
	color := tuple.NewColor(1, 0.8, 0.6)

	for i := range c.Matrix {
		row := c.Matrix[i]
		for j := range row {
			canvas.WritePixel(c, j, i, color)
		}
	}

	ppm := c.ToPPM()

	f, err := os.Create("canvas.ppm")
	if err != nil {
		panic(err)
	}

	f.WriteString(ppm)
}
