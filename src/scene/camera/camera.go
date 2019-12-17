package camera

import (
	"math"

	"github.com/kingsleyliao/ray-tracer/src/rendering/canvas"

	"github.com/kingsleyliao/ray-tracer/src/scene/world"

	"github.com/kingsleyliao/ray-tracer/src/calculation/point"

	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/src/collision/ray"
)

// Camera is the the camera for a scene
type Camera struct {
	Hsize      int
	Vsize      int
	Fov        float64
	Transform  matrix.Matrix
	HalfWidth  float64
	HalfHeight float64
	PixelSize  float64
}

// NewCamera is a constructor for a Camera
func NewCamera(h, v int, fov float64) Camera {
	halfView := math.Tan(fov / 2)
	aspect := float64(h) / float64(v)
	var halfHeight float64
	var halfWidth float64

	if aspect > 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}

	pixelSize := (halfWidth * 2) / float64(h)
	return Camera{
		Hsize:      h,
		Vsize:      v,
		Fov:        fov,
		Transform:  matrix.IdentityMatrix(),
		HalfWidth:  halfWidth,
		HalfHeight: halfHeight,
		PixelSize:  pixelSize,
	}
}

// RayForPixel calculates the Ray that hits a pixel
func (c Camera) RayForPixel(x, y int) ray.Ray {
	xoffset := (float64(x) + 0.5) * c.PixelSize
	yoffset := (float64(y) + 0.5) * c.PixelSize

	worldx := c.HalfWidth - xoffset
	worldy := c.HalfHeight - yoffset

	pixel := c.Transform.Invert().MultiplyTuple(point.NewPoint(worldx, worldy, -1))
	origin := c.Transform.Invert().MultiplyTuple(point.Zero())
	direction := pixel.Subtract(origin).Normalize()
	return ray.NewRay(origin, direction)
}

// Render creates a world
func (c Camera) Render(w world.World) canvas.Canvas {
	image := canvas.NewCanvas(c.Hsize, c.Vsize)

	for y := 0; y < c.Vsize-1; y++ {

		for x := 0; x < c.Hsize-1; x++ {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r)
			canvas.WritePixel(image, x, y, color)
		}
	}

	return image
}
