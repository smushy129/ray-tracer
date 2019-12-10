package light

import (
	"math"

	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
	"github.com/kingsleyliao/ray-tracer/src/rendering/material"
)

// Light in a scene
type Light struct {
	Position  point.Point
	Intensity color.Color
}

// NewPointLight returns a new point light
func NewPointLight(p point.Point, i color.Color) Light {
	return Light{
		Position:  p,
		Intensity: i,
	}
}

// Lighting returns a Color based on the Lighting
func Lighting(m material.Material, l Light, p point.Point, eyeV, normalV vector.Vector) color.Color {
	effectiveColor := m.Color.Multiply(l.Intensity)
	lightVector := l.Position.Subtract(p).Normalize()
	ambient := effectiveColor.Scale(m.Ambient)

	lightDotNormal := lightVector.Dot(normalV)

	var diffuse color.Color
	var specular color.Color
	if lightDotNormal < 0 {
		diffuse = color.Black()
		specular = color.Black()
	} else {
		diffuse = effectiveColor.Scale(m.Diffuse).Scale(lightDotNormal)

		reflectV := lightVector.Invert().Reflect(normalV)
		reflectDotEye := reflectV.Dot(eyeV)

		if reflectDotEye <= 0 {
			specular = color.Black()
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity.Scale(m.Specular * factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
