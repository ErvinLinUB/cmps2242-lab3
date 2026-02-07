package main

import (
	"math"
)

// Part 1: Structs and Methods

// Task 1: Define Struct Types

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Task 2: Implement Methods with Value Receivers

// Area returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the circumference of the circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area returns the area of the triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter returns the perimeter of the triangle (Assuming equilateral triangle: all sides equal to base)
func (t Triangle) Perimeter() float64 {
	return 3 * t.Base
}

// Task 3: Implement Methods with Pointer Receivers

// Scale multiplies the rectangle's dimensions by the given factor
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Scale multiplies the circle's radius by the given factor
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

// Scale multiplies the triangle's dimensions by the given factor
func (t *Triangle) Scale(factor float64) {
	t.Base *= factor
	t.Height *= factor
}
