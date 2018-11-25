package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0
	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct { //creating an "anonymous struct",
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, expected: 100.0},
		{name: "Circle", shape: Circle{10.0}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{10.0, 10.0}, expected: 50.0},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.expected {
			t.Errorf("got %.2f expected %.2f", got, test.expected)
		}
	}
}
