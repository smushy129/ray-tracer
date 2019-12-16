package camera

// Camera is the the camera for a scene
type Camera struct {
	Hsize int
	Vsize int
	Fov   float64
}

// NewCamera is a constructor for a Camera
func NewCamera(h, v int, fov float64) Camera {
	return Camera{
		Hsize: h,
		Vsize: v,
		Fov:   fov,
	}
}
