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
	// Combine the surface color witht he light's color and intensity
	effectiveColor := m.Color.Multiply(l.Intensity)

	// Find the direction of the light source with respect to p (usually zero with unit spheres)
	lightVector := l.Position.Subtract(p).Normalize()

	// Compute the ambient contribution (default ambient is 0.1)
	ambient := effectiveColor.Scale(m.Ambient)

	// lightDotNormal represents the cosine of the angle between the light vector and the normal vector.
	lightDotNormal := lightVector.Dot(normalV)

	var diffuse color.Color
	var specular color.Color
	if lightDotNormal < 0 {
		// A negative number means the light is on the opposite side of the surface
		diffuse = color.Black()
		specular = color.Black()
	} else {
		// compute the diffuse contribution
		diffuse = effectiveColor.Scale(m.Diffuse).Scale(lightDotNormal)
		reflectV := lightVector.Invert().Reflect(normalV)

		// reflectiveDotEye represents the cosine of the angle between the reflection vector and the eye vector
		reflectDotEye := reflectV.Dot(eyeV)

		if reflectDotEye <= 0 {
			// A negative number means the light reflects away from the eye
			specular = color.Black()
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			// Compute the specular contribution
			specular = l.Intensity.Scale(m.Specular * factor)
		}
	}
	// Add all three contributions for the final shading
	return ambient.Add(diffuse).Add(specular)
}
