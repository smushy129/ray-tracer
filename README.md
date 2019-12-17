# Ray Tracer ![build](https://travis-ci.com/kingsleyliao/ray-tracer.svg?branch=master)

This is a ray tracing library capable of rendering 3D objects with realistic lights, shading, and shadows. It contains the primitives and APIs necessary creating a 3D scene including:

- Vectors, Points, Matrices
- Shapes, Lighting, Colors
- Rays, Intersections, Hit testing
- World Scene, Camera

## Requirements:

- [Go Programming language]: https://golang.org
- MacOS (Windows will work, but you will have to find a third party program to open `.ppm` files)

## Steps to run:

1. Go to the root directory and open `main.go`
2. Look at `main()` and decide which renders you want to see outputted - comment out the rest
3. Go to the root directory and run `go build` and then run `./ray-tracer` in your terminal
4. Your render will be output into the project's root directory (change this in `ppm.go` if you like)

**_Note: Some renders will take a while to render depending on your computer specs (5-15 minutes)_**

## Examples

#### Goblin Army

![alt text](https://github.com/kingsleyliao/ray-tracer/blob/master/samples/Goblin%20Army.png "Goblin Army")

#### Purple Giant

![alt text](https://github.com/kingsleyliao/ray-tracer/blob/master/samples/Purple%20Giant.png "Purple Giant")

#### Red Giant

![alt text](https://github.com/kingsleyliao/ray-tracer/blob/master/samples/Red%20Giant.png "Red Giant")

#### Inverted Parabola

![alt text](https://github.com/kingsleyliao/ray-tracer/blob/master/samples/Inverted%20Parabola.png "Inverted Parabola")

#### Clock Face

![alt text](https://github.com/kingsleyliao/ray-tracer/blob/master/samples/Clock%20face.png "Clock face")
