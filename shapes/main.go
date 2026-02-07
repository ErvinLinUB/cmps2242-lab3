package main

import (
	"fmt"
)

//Task 4: Create Main Program

func main() {
	// Create instances of each shape
	rect := Rectangle{Width: 5.0, Height: 3.0}
	circle := Circle{Radius: 2.5}
	triangle := Triangle{Base: 4.0, Height: 3.0}

	fmt.Println("=== Original Shapes ===")

	// Rectangle calculations
	fmt.Printf("Rectangle (Width: %.1f, Height: %.1f):\n", rect.Width, rect.Height)
	fmt.Printf("  Area: %.2f\n", rect.Area())
	fmt.Printf("  Perimeter: %.2f\n\n", rect.Perimeter())

	// Circle calculations
	fmt.Printf("Circle (Radius: %.1f):\n", circle.Radius)
	fmt.Printf("  Area: %.2f\n", circle.Area())
	fmt.Printf("  Perimeter: %.2f\n\n", circle.Perimeter())

	// Triangle calculations
	fmt.Printf("Triangle (Base: %.1f, Height: %.1f):\n", triangle.Base, triangle.Height)
	fmt.Printf("  Area: %.2f\n", triangle.Area())
	fmt.Printf("  Perimeter: %.2f\n\n", triangle.Perimeter())

	fmt.Println("=== After Scaling ===")

	// Scale the shapes
	rect.Scale(2.0)
	circle.Scale(1.5)
	triangle.Scale(0.5)

	// Rectangle after scaling
	fmt.Printf("Rectangle scaled by 2.0 (Width: %.1f, Height: %.1f):\n", rect.Width, rect.Height)
	fmt.Printf("  Area: %.2f\n", rect.Area())
	fmt.Printf("  Perimeter: %.2f\n\n", rect.Perimeter())

	// Circle after scaling
	fmt.Printf("Circle scaled by 1.5 (Radius: %.1f):\n", circle.Radius)
	fmt.Printf("  Area: %.2f\n", circle.Area())
	fmt.Printf("  Perimeter: %.2f\n\n", circle.Perimeter())

	// Triangle after scaling
	fmt.Printf("Triangle scaled by 0.5 (Base: %.1f, Height: %.1f):\n", triangle.Base, triangle.Height)
	fmt.Printf("  Area: %.2f\n", triangle.Area())
	fmt.Printf("  Perimeter: %.2f\n", triangle.Perimeter())
}
