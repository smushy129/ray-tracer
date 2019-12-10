package material

import "github.com/kingsleyliao/ray-tracer/src/rendering/color"

// Material contains values for lighting
type Material struct {
	Color     color.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

// NewMaterial returns a Material with default values
func NewMaterial() Material {
	return Material{
		Color:     color.NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}
