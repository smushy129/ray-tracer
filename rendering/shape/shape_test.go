package shape

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/calculation/matrix"
)

// Creating a new Sphere has a default Transform of the IdentityMatrix
func TestNewSphere(t *testing.T) {
	s := NewSphere()

	expected := matrix.IdentityMatrix()

	if !s.Transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, s.Transform)
	}
}

func TestTransformSphere(t *testing.T) {
	s := NewSphere()
	m := matrix.TranslationMatrix(2, 3, 4)
	s.SetTransform(m)

	expected := matrix.TranslationMatrix(2, 3, 4)

	if !s.Transform.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, s.Transform)
	}
}
