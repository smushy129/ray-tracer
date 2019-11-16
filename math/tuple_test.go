package math

import (
	"testing"
)

func TestTupleEquals__true1(t *testing.T) {
	t1 := Tuple{1, 2, 3, 1}
	t2 := Tuple{1, 2, 3, 1}

	isEqual := t1.equals(t2)

	if isEqual != true {
		t.Errorf("got %v, expected %v", isEqual, true)
	}
}

func TestTupleEquals__true2(t *testing.T) {
	t1 := Tuple{1, 2, 3.000000001, 1}
	t2 := Tuple{1, 2, 3, 1}

	isEqual := t1.equals(t2)

	if isEqual != true {
		t.Errorf("got %v, expected %v", isEqual, true)
	}
}

func TestTupleEquals__false(t *testing.T) {
	t1 := Tuple{1, 2, 3, 0}
	t2 := Tuple{1, 2, 3, 1}

	isEqual := t1.equals(t2)

	if isEqual != false {
		t.Errorf("got %v, expected %v", isEqual, false)
	}
}

func TestTupleAdd(t *testing.T) {
	t1 := Tuple{3, -2, 5, 1}
	t2 := Tuple{-2, 3, 1, 0}

	sum := t1.add(t2)
	expected := Tuple{1, 1, 6, 1}

	if !sum.equals(expected) {
		t.Errorf("got %v, expected %v", sum, expected)
	}
}

func TestTupleInvert(t *testing.T) {
	t1 := Tuple{1, 2, 3, -4}

	inverted := t1.invert()
	expected := Tuple{-1, -2, -3, 4}

	if !inverted.equals(expected) {
		t.Errorf("got %v, expected %v", inverted, expected)
	}
}

func TestTupleScale(t *testing.T) {
	t1 := Tuple{1, 2, 3, -4}

	scaled := t1.scale(2)
	expected := Tuple{2, 4, 6, -8}

	if !scaled.equals(expected) {
		t.Errorf("got %v, expected %v", scaled, expected)
	}
}
