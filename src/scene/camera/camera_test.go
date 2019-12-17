package camera

import (
	"math"
	"testing"

	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/rendering/canvas"
	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
	"github.com/kingsleyliao/ray-tracer/src/scene/world"

	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"

	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
)

// Construting a Camera
func TestNewCamera(t *testing.T) {
	hsize := 160
	vsize := 120
	fov := math.Pi / 2

	c := NewCamera(hsize, vsize, fov)

	if c.Hsize != 160 {
		t.Errorf("expected %v, got %v", 160, c.Hsize)
	}

	if c.Vsize != 120 {
		t.Errorf("expected %v, got %v", 120, c.Vsize)
	}

	if c.Fov != math.Pi/2 {
		t.Errorf("expected %v, got %v", math.Pi/2, c.Fov)
	}

	if !c.Transform.Equals(matrix.IdentityMatrix()) {
		t.Errorf("expected %v, got %v", matrix.IdentityMatrix(), c.Transform)
	}
}

// The pixel size of a horizontal camera
func TestHorizontalCamera(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2)
	expected := 0.01

	if c.PixelSize != expected {
		t.Errorf("expected %v, got %v", expected, c.PixelSize)
	}
}

// The pixel size of a vertical camera
func TestVerticalCamera(t *testing.T) {
	c := NewCamera(125, 200, math.Pi/2)
	expected := 0.01

	if c.PixelSize != expected {
		t.Errorf("expected %v, got %v", expected, c.PixelSize)
	}
}

// Constructing a ray through the center of the canvas
func TestRayThroughCanvas_1(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)

	expectedOrigin := point.Zero()
	expectedDirection := vector.Front()

	if !r.Origin.Equals(expectedOrigin) {
		t.Errorf("expected %v, got %v", expectedOrigin, r.Origin)
	}

	if !r.Direction.Equals(expectedDirection) {
		t.Errorf("expected %v, got %v", expectedDirection, r.Direction)
	}
}

// Constructing a ray through the corner of the canvas
func TestRayThroughCanvas_2(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(0, 0)

	expectedOrigin := point.Zero()
	expectedDirection := vector.NewVector(0.66519, 0.33259, -0.66851)

	if !r.Origin.Equals(expectedOrigin) {
		t.Errorf("expected %v, got %v", expectedOrigin, r.Origin)
	}

	if !r.Direction.Equals(expectedDirection) {
		t.Errorf("expected %v, got %v", expectedDirection, r.Direction)
	}
}

// Constructing a ray when the camera is transformed
func TestRayThroughCanvas_3(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	c.Transform = matrix.RotationMatrix(vector.Up(), math.Pi/4).Multiply(matrix.TranslationMatrix(0, -2, 5))
	r := c.RayForPixel(100, 50)

	expectedOrigin := point.NewPoint(0, 2, -5)
	expectedDirection := vector.NewVector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2)

	if !r.Origin.Equals(expectedOrigin) {
		t.Errorf("expected %v, got %v", expectedOrigin, r.Origin)
	}

	if !r.Direction.Equals(expectedDirection) {
		t.Errorf("expected %v, got %v", expectedDirection, r.Direction)
	}
}

// Rendering a world with a camera
func TestRenderWorldWithCamera(t *testing.T) {
	w := world.DefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)
	from := point.NewPoint(0, 0, -5)
	to := point.NewPoint(0, 0, 0)
	up := vector.Up()
	c.Transform = world.ViewTransform(from, to, up)
	image := c.Render(w)
	pixelColor, _ := canvas.PixelAt(image, 5, 5)

	expected := color.NewColor(0.38066, 0.47583, 0.2855)

	if !pixelColor.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, pixelColor)
	}
}
