package math

import "math"

// Tuple represents Points and Vectors
type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func (t Tuple) equals(t1 Tuple) bool {
	EPSILON := 0.00001
	if math.Abs(t.x-t1.x) > EPSILON {
		return false
	}
	if math.Abs(t.y-t1.y) > EPSILON {
		return false
	}
	if math.Abs(t.z-t1.z) > EPSILON {
		return false
	}
	if t.w != t1.w {
		return false
	}
	return true
}

func (t Tuple) add(t1 Tuple) Tuple {
	return Tuple{
		x: t.x + t1.x,
		y: t.y + t1.y,
		z: t.z + t1.z,
		w: t.w + t1.w,
	}
}

func (t Tuple) subtract(t1 Tuple) Tuple {
	return Tuple{
		x: t.x - t1.x,
		y: t.y - t1.y,
		z: t.z - t1.z,
		w: t.w - t1.w,
	}
}

func (t Tuple) scale(s float64) Tuple {
	return Tuple{
		x: t.x * s,
		y: t.y * s,
		z: t.z * s,
		w: t.w * s,
	}
}

func (t Tuple) invert() Tuple {
	return Tuple{
		x: -t.x,
		y: -t.y,
		z: -t.z,
		w: -t.w,
	}
}
