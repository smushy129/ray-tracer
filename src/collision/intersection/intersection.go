package intersection

import (
	"math"
	"sort"

	"github.com/kingsleyliao/ray-tracer/src/collision/ray"
	"github.com/kingsleyliao/ray-tracer/src/env/world"
	"github.com/kingsleyliao/ray-tracer/src/rendering/shape"
)

// Intersection encapsulates t and object
type Intersection struct {
	T      float64
	Object shape.Sphere
}

// Equals compares Intersections for equality
func (i Intersection) Equals(i2 Intersection) bool {
	return i.T == i2.T &&
		i.Object.Equals(i2.Object)
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
	r2 := r.Transform(s.Transform.Invert())
	sphereToRay := r2.Origin.Subtract(s.Center)

	a := r2.Direction.Dot(r2.Direction)
	b := 2.0 * r2.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - (4 * a * c)
	if discriminant < 0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

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

// IntersectWorld returns a all the objects in a world intersected by a Ray
func IntersectWorld(w world.World, r ray.Ray) ByT {
	intersections := ByT{}
	for _, s := range w.Objects {
		intersections = append(intersections, Intersect(s, r)...)
	}
	sort.Sort(intersections)
	return intersections
}
