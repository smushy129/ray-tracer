package ray

import (
	"testing"

	"github.com/kingsleyliao/ray-tracer/calculation/point"
	"github.com/kingsleyliao/ray-tracer/calculation/vector"
)

func TestNewRay(t *testing.T) {
	origin := point.NewPoint(1, 2, 3)
	direction := vector.NewVector(4, 5, 6)

	ray := NewRay(origin, direction)
	result := ray.Direction == direction && ray.Origin == origin
	expected := true

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPositionAt_1(t *testing.T) {
	ray := NewRay(point.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0))

	position := ray.PositionAt(0)
	expected := point.NewPoint(2, 3, 4)

	if !position.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, position)
	}
}

func TestPositionAt_2(t *testing.T) {
	ray := NewRay(point.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0))

	position := ray.PositionAt(1)
	expected := point.NewPoint(3, 3, 4)

	if !position.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, position)
	}
}

func TestPositionAt_3(t *testing.T) {
	ray := NewRay(point.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0))

	position := ray.PositionAt(-1)
	expected := point.NewPoint(1, 3, 4)

	if !position.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, position)
	}
}

func TestPositionAt_4(t *testing.T) {
	ray := NewRay(point.NewPoint(2, 3, 4), vector.NewVector(1, 0, 0))

	position := ray.PositionAt(2.5)
	expected := point.NewPoint(4.5, 3, 4)

	if !position.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, position)
	}
}
