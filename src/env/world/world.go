package world

import (
	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
	"github.com/kingsleyliao/ray-tracer/src/rendering/light"
	"github.com/kingsleyliao/ray-tracer/src/rendering/material"
	"github.com/kingsleyliao/ray-tracer/src/rendering/shape"
)

// World represents the entire 3D scene
type World struct {
	Objects []shape.Sphere
	Light   light.Light
}

// DefaultWorld creates a World with 1 point light and 2 Spheres
func DefaultWorld() World {
	s1 := shape.NewSphere()
	s1.Material = material.Material{
		Color:    color.NewColor(0.8, 1.0, 0.6),
		Diffuse:  0.7,
		Specular: 0.2,
	}
	s2 := shape.NewSphere()
	s2.Transform = matrix.ScalingMatrix(0.5, 0.5, 0.5)

	l := light.NewPointLight(point.NewPoint(-10, 10, -10), color.NewColor(1, 1, 1))

	return World{
		Objects: []shape.Sphere{s1, s2},
		Light:   l,
	}
}
