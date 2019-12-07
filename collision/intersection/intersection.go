package intersection

import (
	"math"

	"github.com/kingsleyliao/ray-tracer/collision/ray"
	"github.com/kingsleyliao/ray-tracer/rendering/shape"
)

// Intersect finds the intersection points (time T) of a ray and a sphere
func Intersect(s shape.Sphere, r ray.Ray) []float64 {
	sphereToRay := r.Origin.Subtract(s.Center)

	a := r.Direction.Dot(r.Direction)
	b := 2.0 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discrminant := (b * b) - (4 * a * c)
	if discrminant < 0 {
		return nil
	}

	t1 := (-b - math.Sqrt(discrminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discrminant)) / (2 * a)
	return []float64{t1, t2}
}
