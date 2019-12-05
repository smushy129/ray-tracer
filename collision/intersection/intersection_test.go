package intersection

import (
	"testing"

	sphere "github.com/kingsleyliao/ray-tracer/rendering/shape"

	"github.com/kingsleyliao/ray-tracer/calculation/point"
	"github.com/kingsleyliao/ray-tracer/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/collision/ray"
)

func TestIntersect(t *testing.T) {
	r := ray.NewRay(point.NewPoint(0, 0, -5), vector.NewVector(0, 0, 1))
	s := sphere.NewSphere()

	xs := Intersect(s, r)

	expected0 := 4.0 == xs[0]
	expected1 := 6.0 == xs[1]

	if expected0 != true || expected1 != true {
		t.Errorf("expected %v, %v, got %v, %v", expected0, expected1, xs[0], xs[1])
	}
}
