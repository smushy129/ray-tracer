package renders

import (
	"math"

	"github.com/kingsleyliao/ray-tracer/src/calculation/point"
	"github.com/kingsleyliao/ray-tracer/src/rendering/light"
	"github.com/kingsleyliao/ray-tracer/src/scene/camera"
	"github.com/kingsleyliao/ray-tracer/src/scene/world"

	"github.com/kingsleyliao/ray-tracer/src/calculation/matrix"
	"github.com/kingsleyliao/ray-tracer/src/calculation/vector"
	"github.com/kingsleyliao/ray-tracer/src/rendering/color"
	"github.com/kingsleyliao/ray-tracer/src/rendering/material"
	"github.com/kingsleyliao/ray-tracer/src/rendering/shape"
)

// DrawGoblinArmy draws the goblin army scene
func DrawGoblinArmy() {
	floor := shape.NewSphere()
	floor.Transform = matrix.ScalingMatrix(10, 0.01, 10)
	floor.Material = material.DefaultMaterial()
	floor.Material.Color = color.NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0

	leftWall := shape.NewSphere()
	leftWall.Transform = matrix.TranslationMatrix(0, 0, 5).
		Multiply(matrix.RotationMatrix(vector.Up(), -math.Pi/4)).
		Multiply(matrix.RotationMatrix(vector.Right(), math.Pi/2)).
		Multiply(matrix.ScalingMatrix(10, 0.01, 10))
	leftWall.Material = floor.Material

	rightWall := shape.NewSphere()
	rightWall.Transform = matrix.TranslationMatrix(0, 0, 5).
		Multiply(matrix.RotationMatrix(vector.Up(), math.Pi/4)).
		Multiply(matrix.RotationMatrix(vector.Right(), math.Pi/2)).
		Multiply(matrix.ScalingMatrix(10, 0.01, 10))
	rightWall.Material = floor.Material

	bigSphere := shape.NewSphere()
	bigSphere.Transform = matrix.TranslationMatrix(-0.5, 1, 0.5)
	bigSphere.Material = material.DefaultMaterial()
	bigSphere.Material.Color = color.NewColor(0.1, 1, 0.5)
	bigSphere.Material.Diffuse = 0.7
	bigSphere.Material.Specular = 0.3

	mediumSphere := shape.NewSphere()
	mediumSphere.Transform = matrix.TranslationMatrix(1.5, 0.5, -0.5).Multiply(matrix.ScalingMatrix(0.5, 0.5, 0.5))
	mediumSphere.Material = material.DefaultMaterial()
	mediumSphere.Material.Color = color.NewColor(0.1, 1, 0.5)
	mediumSphere.Material.Diffuse = 0.7
	mediumSphere.Material.Specular = 0.3

	tanSphere := shape.NewSphere()
	tanSphere.Transform = matrix.TranslationMatrix(-1.5, 0.33, -0.75).Multiply(matrix.ScalingMatrix(0.33, 0.33, 0.33))
	tanSphere.Material = material.DefaultMaterial()
	tanSphere.Material.Color = color.NewColor(1, 0.8, 0.1)
	tanSphere.Material.Diffuse = 0.7
	tanSphere.Material.Specular = 0.3

	w := world.World{
		Light:   light.NewPointLight(point.NewPoint(-10, 10, -10), color.NewColor(1, 1, 1)),
		Objects: []shape.Sphere{tanSphere, bigSphere, mediumSphere, floor, leftWall, rightWall},
	}

	camera := camera.NewCamera(500, 500, math.Pi/3)
	camera.Transform = world.ViewTransform(point.NewPoint(0, 1.5, -5), point.NewPoint(0, 1, 0), vector.Up())

	canvas := camera.Render(w)
	ppm := canvas.ToPPM()
	outputPPM(ppm, "goblinarmy.ppm")
}
