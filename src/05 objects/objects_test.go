package objects

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Version 3: Test Perimeter with a structure
func TestPerimeter3(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter3(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Version 3: Test Area with a structure
func TestArea3(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Area3(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
