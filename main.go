package main

import (
	"fmt"

	m "github.com/kingsleyliao/ray-tracer/math"
)

func main() {
	v := m.NewVector(1, 2, 3)
	p := m.NewPoint(1, 2, 3)

	v.Magnitude()
	p.Magnitude()

	fmt.Printf("%v \n", v.Magnitude())
	fmt.Printf("%v \n", p.Magnitude())
	fmt.Printf("%v \n", v.add(p))
}
