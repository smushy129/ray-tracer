package projectile

import (
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
)

// Environment is the env a projectile is projected in
type Environment interface {
	Gravity() vector.Vector
	Wind() vector.Vector
}

type Earth struct {
	gravity vector.Vector
	wind    vector.Vector
}

func NewEarth(gravity, wind vector.Vector) Earth {
	return Earth{gravity, wind}
}

func (e Earth) Wind() vector.Vector {
	return e.wind
}

func (e Earth) Gravity() vector.Vector {
	return e.gravity
}

// Projectile is an object that travels through space on each tick
type Projectile interface {
	Tick(e Environment)
}

type Ray struct {
	Point    point.Point
	Velocity vector.Vector
}

func NewRay(p point.Point, v vector.Vector) Ray {
	return Ray{p, v}
}

func (r Ray) Tick(e Environment) Ray {
	position := r.Point.Add(r.Velocity)
	velocity := r.Velocity.Add(e.Gravity()).Add(e.Wind())
	return Ray{position, velocity}
}
