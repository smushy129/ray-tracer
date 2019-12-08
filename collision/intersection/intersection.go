package intersection

import (
	"math"
	"sort"

	"github.com/kingsleyliao/ray-tracer/collision/ray"
	"github.com/kingsleyliao/ray-tracer/rendering/shape"
)

// Intersection encapsulates t and object
type Intersection struct {
	T      float64
	Object shape.Sphere
}

// ByT is used to sort Intersections by T value
type ByT []Intersection

func (xs ByT) Len() int           { return len(xs) }
func (xs ByT) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs ByT) Less(i, j int) bool { return xs[i].T < xs[j].T }

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

// Hit returns the first Intersection of a ray
func Hit(xs []Intersection) (x Intersection, ok bool) {
	sort.Sort(ByT(xs))
	hit := Intersection{ // default value
		T:      -1, // ok will be false if T is negative
		Object: shape.NewSphere(),
	}

	// Find the lowest non=negative intersection
	for _, intersection := range xs {
		if intersection.T >= 0 {
			hit = intersection
			break
		}
	}

	// If all values of T are < 0, no object was hit
	if hit.T < 0 {
		return hit, false
	}
	return hit, true
}
