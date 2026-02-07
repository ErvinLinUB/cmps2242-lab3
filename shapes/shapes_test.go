package main

import (
	"math"
	"testing"
)

// Task 5: Write Unit Tests

// TestRectangleArea tests the Area() method for Rectangle
func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	expected := 15.0
	actual := rect.Area()

	if math.Abs(actual-expected) > 0.0001 {
		t.Errorf("Rectangle.Area() = %v, expected %v", actual, expected)
	}
}

// TestRectanglePerimeter tests the Perimeter() method for Rectangle
func TestRectanglePerimeter(t *testing.T) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	expected := 16.0
	actual := rect.Perimeter()

	if math.Abs(actual-expected) > 0.0001 {
		t.Errorf("Rectangle.Perimeter() = %v, expected %v", actual, expected)
	}
}

// TestRectangleScale tests the Scale() method for Rectangle
func TestRectangleScale(t *testing.T) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	rect.Scale(2.0)

	expectedWidth := 10.0
	expectedHeight := 6.0

	if math.Abs(rect.Width-expectedWidth) > 0.0001 {
		t.Errorf("After Scale(2.0): Rectangle.Width = %v, expected %v", rect.Width, expectedWidth)
	}

	if math.Abs(rect.Height-expectedHeight) > 0.0001 {
		t.Errorf("After Scale(2.0): Rectangle.Height = %v, expected %v", rect.Height, expectedHeight)
	}

	// Also test that area is updated correctly
	expectedArea := 60.0
	actualArea := rect.Area()
	if math.Abs(actualArea-expectedArea) > 0.0001 {
		t.Errorf("After scaling, Rectangle.Area() = %v, expected %v", actualArea, expectedArea)
	}
}

// TestCircleArea tests the Area() method for Circle
func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 2.5}
	expected := math.Pi * 2.5 * 2.5
	actual := circle.Area()

	if math.Abs(actual-expected) > 0.0001 {
		t.Errorf("Circle.Area() = %v, expected %v", actual, expected)
	}
}

// TestCirclePerimeter tests the Perimeter() method for Circle
func TestCirclePerimeter(t *testing.T) {
	circle := Circle{Radius: 2.5}
	expected := 2 * math.Pi * 2.5
	actual := circle.Perimeter()

	if math.Abs(actual-expected) > 0.0001 {
		t.Errorf("Circle.Perimeter() = %v, expected %v", actual, expected)
	}
}

// TestCircleScale tests the Scale() method for Circle
func TestCircleScale(t *testing.T) {
	circle := Circle{Radius: 2.5}
	circle.Scale(1.5)

	expectedRadius := 3.75

	if math.Abs(circle.Radius-expectedRadius) > 0.0001 {
		t.Errorf("After Scale(1.5): Circle.Radius = %v, expected %v", circle.Radius, expectedRadius)
	}

	// Also test that area is updated correctly
	expectedArea := math.Pi * expectedRadius * expectedRadius
	actualArea := circle.Area()
	if math.Abs(actualArea-expectedArea) > 0.0001 {
		t.Errorf("After scaling, Circle.Area() = %v, expected %v", actualArea, expectedArea)
	}
}

// TestTriangleArea tests the Area() method for Triangle
func TestTriangleArea(t *testing.T) {
	triangle := Triangle{Base: 4.0, Height: 3.0}
	expected := 6.0
	actual := triangle.Area()

	if math.Abs(actual-expected) > 0.0001 {
		t.Errorf("Triangle.Area() = %v, expected %v", actual, expected)
	}
}

// TestTrianglePerimeter tests the Perimeter() method for Triangle
func TestTrianglePerimeter(t *testing.T) {
	triangle := Triangle{Base: 4.0, Height: 3.0}
	expected := 12.0 // 3 * base (assuming equilateral triangle)
	actual := triangle.Perimeter()

	if math.Abs(actual-expected) > 0.0001 {
		t.Errorf("Triangle.Perimeter() = %v, expected %v", actual, expected)
	}
}

// TestTriangleScale tests the Scale() method for Triangle
func TestTriangleScale(t *testing.T) {
	triangle := Triangle{Base: 4.0, Height: 3.0}
	triangle.Scale(0.5)

	expectedBase := 2.0
	expectedHeight := 1.5

	if math.Abs(triangle.Base-expectedBase) > 0.0001 {
		t.Errorf("After Scale(0.5): Triangle.Base = %v, expected %v", triangle.Base, expectedBase)
	}

	if math.Abs(triangle.Height-expectedHeight) > 0.0001 {
		t.Errorf("After Scale(0.5): Triangle.Height = %v, expected %v", triangle.Height, expectedHeight)
	}

	// Also test that area is updated correctly
	expectedArea := 1.5 // 0.5 * 2.0 * 1.5
	actualArea := triangle.Area()
	if math.Abs(actualArea-expectedArea) > 0.0001 {
		t.Errorf("After scaling, Triangle.Area() = %v, expected %v", actualArea, expectedArea)
	}
}
