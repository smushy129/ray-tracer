package projectile

import (
	"github.com/kingsleyliao/ray-tracer/tuple"
)

// Environment is the env a projectile is projected in
type Environment interface {
	Gravity() tuple.Vector
	Wind() tuple.Vector
}

type Earth struct {
	gravity tuple.Vector
	wind    tuple.Vector
}

func NewEarth(gravity, wind tuple.Vector) Earth {
	return Earth{gravity, wind}
}

func (e Earth) Wind() tuple.Vector {
	return e.wind
}

func (e Earth) Gravity() tuple.Vector {
	return e.gravity
}

// Projectile is an object that travels through space on each tick
type Projectile interface {
	Tick(e Environment)
}

type Ray struct {
	Point    tuple.Point
	Velocity tuple.Vector
}

func NewRay(p tuple.Point, v tuple.Vector) Ray {
	return Ray{p, v}
}

func (r Ray) Tick(e Environment) Ray {
	position := r.Point.Add(r.Velocity)
	velocity := r.Velocity.Add(e.Gravity()).Add(e.Wind())
	return Ray{position, velocity}
}
