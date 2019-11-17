package canvas

import (
	"fmt"
	"testing"

	"github.com/kingsleyliao/ray-tracer/tuple"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)

	expected := Canvas{
		Width:  10,
		Height: 20,
	}

	if c.Width != expected.Width ||
		c.Height != expected.Height {
		t.Errorf("expected %v, got %v", expected, c)
	}
}

func TestWritePixel(t *testing.T) {
	c := NewCanvas(10, 20)
	red := tuple.NewColor(1, 0, 0)

	WritePixel(c, 2, 3, red)
	pixel, err := PixelAt(c, 2, 3)

	if err != nil {
		fmt.Println(err)
	}
	if pixel != red {
		t.Errorf("expected %v, got %v", red, pixel)
	}
}

func TestToPPM(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := c.ToPPM()
	expected :=
		"P3\n" + "5 3\n" + "255\n" +
			"0 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n"

	if ppm != expected {
		t.Errorf("expected:\n %v,\n got:\n %v", expected, ppm)
	}

}

func TestToPPM_2(t *testing.T) {
	c := NewCanvas(5, 3)
	c1 := tuple.NewColor(1.5, 0, 0)
	c2 := tuple.NewColor(0, 0.5, 0)
	c3 := tuple.NewColor(-0.5, 0, 1)

	WritePixel(c, 0, 0, c1)
	WritePixel(c, 2, 1, c2)
	WritePixel(c, 4, 2, c3)

	ppm := c.ToPPM()
	expected :=
		"P3\n" + "5 3\n" + "255\n" +
			"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n" +
			"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n"

	if ppm != expected {
		t.Errorf("expected:\n %v,\n got:\n %v", expected, ppm)
	}
}

func TestToPPM_3(t *testing.T) {
	c := NewCanvas(10, 2)
	color := tuple.NewColor(1, 0.8, 0.6)

	for i := range c.Matrix {
		row := c.Matrix[i]
		for j := range row {
			WritePixel(c, j, i, color)
		}
	}

	ppm := c.ToPPM()
	expected :=
		"P3\n" + "10 2\n" + "255\n" +
			"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153\n" +
			"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153\n"

	if ppm != expected {
		t.Errorf("expected:\n %v,\n got:\n %v", expected, ppm)
	}
}
