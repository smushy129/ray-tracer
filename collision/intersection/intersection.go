package intersection

import (
	"math"

	"github.com/kingsleyliao/ray-tracer/collision/ray"
	"github.com/kingsleyliao/ray-tracer/rendering/shape"
)

// Intersection encapsulates t and object
type Intersection struct {
	T      float64
	Object shape.Sphere
}

// NewIntersection is the constructor for Intersection
func NewIntersection(t float64, s shape.Sphere) Intersection {
	return Intersection{
		T:      t,
		Object: s,
	}
}

// Intersections creates a slice of Intersection
func Intersections(xs ...Intersection) []Intersection {
	intersections := make([]Intersection, 0)
	for _, x := range xs {
		intersections = append(intersections, x)
	}
	return intersections
}

// Intersect finds the intersection points (time T) of a ray and a sphere
func Intersect(s shape.Sphere, r ray.Ray) []Intersection {
	sphereToRay := r.Origin.Subtract(s.Center)

	a := r.Direction.Dot(r.Direction)
	b := 2.0 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discrminant := (b * b) - (4 * a * c)
	if discrminant < 0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(discrminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discrminant)) / (2 * a)

	i1 := Intersection{T: t1, Object: s}
	i2 := Intersection{T: t2, Object: s}

	return []Intersection{i1, i2}
}
